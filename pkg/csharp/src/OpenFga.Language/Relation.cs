using System.Collections.Generic;
using OpenFga.Sdk.Model;

namespace OpenFga.Language
{
    /// <summary>
    /// Helper class to hold relation information during parsing
    /// This bridges the gap between the ANTLR parsing and the OpenFGA SDK model classes
    /// </summary>
    public class Relation
    {
        public List<Userset> Rewrites { get; set; }
        public string Operator { get; set; }
        public RelationMetadata TypeInfo { get; set; }

        public Relation()
        {
            Rewrites = new List<Userset>();
            TypeInfo = new RelationMetadata();
        }

        public Relation(object unused1, List<Userset> rewrites, object unused2, RelationMetadata typeInfo)
        {
            Rewrites = rewrites ?? new List<Userset>();
            TypeInfo = typeInfo ?? new RelationMetadata();
        }
    }
}
