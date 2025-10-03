using System.Collections.Generic;
using OpenFga.Sdk.Model;

namespace OpenFga.Language
{
    /// <summary>
    /// Helper class to store relation information on the stack during parsing
    /// </summary>
    public class StackRelation
    {
        public List<Userset> Rewrites { get; set; }
        public string Operator { get; set; }

        public StackRelation(List<Userset> rewrites, string @operator)
        {
            Rewrites = rewrites ?? new List<Userset>();
            Operator = @operator;
        }
    }
}
