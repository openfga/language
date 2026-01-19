package transformer

import (
	"strings"
)

// CommentInfo holds comment information for a DSL element.
type CommentInfo struct {
	// PrecedingLines contains comment lines that appear before the element (including the # prefix)
	PrecedingLines []string `json:"preceding_lines,omitempty"`
	// Inline contains an inline comment that appears on the same line as the element (including the # prefix)
	Inline string `json:"inline,omitempty"`
}

// CommentsMetadata holds comments for types, relations, and conditions.
type CommentsMetadata struct {
	// Comments for the element itself
	Comments *CommentInfo `json:"comments,omitempty"`
	// RelationComments maps relation names to their comments
	RelationComments map[string]*CommentInfo `json:"relation_comments,omitempty"`
}

// ModelComments holds comments at the model level.
type ModelComments struct {
	// PrecedingLines contains comment lines that appear before the model declaration
	PrecedingLines []string `json:"preceding_lines,omitempty"`
}

// CommentTracker tracks comments in DSL source and maps them to AST elements.
type CommentTracker struct {
	lines            []string
	lineComments     map[int]string           // line number -> inline comment
	modelComment     *ModelComments           // comments before the model declaration
	typeComments     map[string]*CommentsMetadata  // type name -> comments metadata
	condComments     map[string]*CommentInfo       // condition name -> comment info
	cleanedToOriginal map[int]int              // mapping from cleaned line numbers to original
}

// NewCommentTracker creates a new CommentTracker from DSL source.
func NewCommentTracker(source string) *CommentTracker {
	ct := &CommentTracker{
		lines:            strings.Split(source, "\n"),
		lineComments:     make(map[int]string),
		typeComments:     make(map[string]*CommentsMetadata),
		condComments:     make(map[string]*CommentInfo),
		cleanedToOriginal: nil,
	}
	ct.parseComments()
	return ct
}

// NewCommentTrackerWithMapping creates a new CommentTracker with line number mapping.
func NewCommentTrackerWithMapping(source string, cleanedToOriginal map[int]int) *CommentTracker {
	ct := &CommentTracker{
		lines:            strings.Split(source, "\n"),
		lineComments:     make(map[int]string),
		typeComments:     make(map[string]*CommentsMetadata),
		condComments:     make(map[string]*CommentInfo),
		cleanedToOriginal: cleanedToOriginal,
	}
	ct.parseComments()
	return ct
}

// parseComments parses all comments from the source.
func (ct *CommentTracker) parseComments() {
	for i, line := range ct.lines {
		// Check for inline comment
		if inlineIdx := strings.Index(line, " #"); inlineIdx != -1 {
			// Make sure this isn't inside a condition expression or similar
			// We only consider it an inline comment if there's actual code before it
			beforeComment := strings.TrimSpace(line[:inlineIdx])
			if len(beforeComment) > 0 {
				ct.lineComments[i] = strings.TrimSpace(line[inlineIdx+1:])
			}
		}
	}
}

// GetPrecedingComments returns comments that immediately precede the given line number.
// Line numbers are 0-based.
func (ct *CommentTracker) GetPrecedingComments(lineNum int) []string {
	if lineNum <= 0 || lineNum > len(ct.lines) {
		return nil
	}

	var comments []string
	// Walk backwards from the line before the target
	for i := lineNum - 1; i >= 0; i-- {
		line := strings.TrimSpace(ct.lines[i])
		if len(line) == 0 {
			// Empty line breaks the contiguous comment block
			break
		}
		if strings.HasPrefix(line, "#") {
			// Prepend to maintain order
			comments = append([]string{line}, comments...)
		} else {
			// Non-comment, non-empty line breaks the block
			break
		}
	}

	return comments
}

// GetInlineComment returns the inline comment for the given line number.
// Line numbers are 0-based.
func (ct *CommentTracker) GetInlineComment(lineNum int) string {
	if lineNum < 0 || lineNum >= len(ct.lines) {
		return ""
	}

	line := ct.lines[lineNum]
	// Find inline comment
	if inlineIdx := strings.Index(line, " #"); inlineIdx != -1 {
		beforeComment := strings.TrimSpace(line[:inlineIdx])
		if len(beforeComment) > 0 {
			return strings.TrimSpace(line[inlineIdx+1:])
		}
	}
	return ""
}

// GetModelComments returns comments that appear before the model declaration.
func (ct *CommentTracker) GetModelComments() *ModelComments {
	// Find the model declaration line
	modelLine := -1
	for i, line := range ct.lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "model") {
			modelLine = i
			break
		}
	}

	if modelLine <= 0 {
		return nil
	}

	precedingComments := ct.GetPrecedingComments(modelLine)
	if len(precedingComments) == 0 {
		return nil
	}

	return &ModelComments{
		PrecedingLines: precedingComments,
	}
}

// GetCommentInfoForLine creates a CommentInfo for elements at the given line.
// lineNum is expected to be 0-based and may be from cleaned source (if mapping exists).
func (ct *CommentTracker) GetCommentInfoForLine(lineNum int) *CommentInfo {
	// If we have a line mapping, convert from cleaned line number to original
	originalLineNum := lineNum
	if ct.cleanedToOriginal != nil {
		if mapped, ok := ct.cleanedToOriginal[lineNum]; ok {
			originalLineNum = mapped
		}
	}

	preceding := ct.GetPrecedingComments(originalLineNum)
	inline := ct.GetInlineComment(originalLineNum)

	if len(preceding) == 0 && inline == "" {
		return nil
	}

	return &CommentInfo{
		PrecedingLines: preceding,
		Inline:         inline,
	}
}

// SetTypeComments sets comments for a type at the given line.
func (ct *CommentTracker) SetTypeComments(typeName string, lineNum int) {
	commentInfo := ct.GetCommentInfoForLine(lineNum)
	if commentInfo == nil {
		return
	}

	if ct.typeComments[typeName] == nil {
		ct.typeComments[typeName] = &CommentsMetadata{}
	}
	ct.typeComments[typeName].Comments = commentInfo
}

// SetRelationComments sets comments for a relation within a type.
func (ct *CommentTracker) SetRelationComments(typeName, relationName string, lineNum int) {
	commentInfo := ct.GetCommentInfoForLine(lineNum)
	if commentInfo == nil {
		return
	}

	if ct.typeComments[typeName] == nil {
		ct.typeComments[typeName] = &CommentsMetadata{}
	}
	if ct.typeComments[typeName].RelationComments == nil {
		ct.typeComments[typeName].RelationComments = make(map[string]*CommentInfo)
	}
	ct.typeComments[typeName].RelationComments[relationName] = commentInfo
}

// SetConditionComments sets comments for a condition.
func (ct *CommentTracker) SetConditionComments(conditionName string, lineNum int) {
	commentInfo := ct.GetCommentInfoForLine(lineNum)
	if commentInfo == nil {
		return
	}
	ct.condComments[conditionName] = commentInfo
}

// GetTypeComments returns the comments metadata for a type.
func (ct *CommentTracker) GetTypeComments(typeName string) *CommentsMetadata {
	return ct.typeComments[typeName]
}

// GetRelationComments returns comments for a relation.
func (ct *CommentTracker) GetRelationComments(typeName, relationName string) *CommentInfo {
	if ct.typeComments[typeName] == nil {
		return nil
	}
	return ct.typeComments[typeName].RelationComments[relationName]
}

// GetConditionComments returns comments for a condition.
func (ct *CommentTracker) GetConditionComments(conditionName string) *CommentInfo {
	return ct.condComments[conditionName]
}
