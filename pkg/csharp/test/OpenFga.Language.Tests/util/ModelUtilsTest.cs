using OpenFga.Language.utils;
using OpenFga.Sdk.Model;
using Xunit;

namespace OpenFga.Language.Tests.util;

public class ModelUtilsTest {
    private TypeDefinition GetTypeDefWithModules() {
        return new TypeDefinition {
            Type = "type1",
            Relations = new Dictionary<string, Userset>
            {
                { "relation1", new Userset() },
                { "relation2", new Userset() },
                { "relation3", new Userset() },
                { "relation4", new Userset() }
            },
            Metadata = new Metadata {
                Module = "type_module1",
                Relations = new Dictionary<string, RelationMetadata>
                {
                    { "relation1", new RelationMetadata { Module = "module1" } },
                    { "relation2", new RelationMetadata { Module = "" } },
                    { "relation3", new RelationMetadata() },
                    { "relation5", new RelationMetadata() }
                }
            }
        };
    }

    private TypeDefinition GetTypeDefWithoutModules() {
        return new TypeDefinition {
            Type = "type2",
            Relations = new Dictionary<string, Userset>
            {
                { "relation7", new Userset() }
            }
        };
    }

    [Fact]
    public void TestGetModuleForObjectTypeRelation_RelationExistsAndHasModule() {
        var result = ModelUtils.GetModuleForObjectTypeRelation(GetTypeDefWithModules(), "relation1");
        Assert.Equal("module1", result);
    }

    [Fact]
    public void TestGetModuleForObjectTypeRelation_RelationExistsButHasEmptyModule_TypeHasModule() {
        var result = ModelUtils.GetModuleForObjectTypeRelation(GetTypeDefWithModules(), "relation2");
        Assert.Equal("type_module1", result);
    }

    [Fact]
    public void TestGetModuleForObjectTypeRelation_RelationExistsButNoModule_TypeHasModule() {
        var result = ModelUtils.GetModuleForObjectTypeRelation(GetTypeDefWithModules(), "relation3");
        Assert.Equal("type_module1", result);
    }

    [Fact]
    public void TestGetModuleForObjectTypeRelation_RelationExistsButNoMetadata_TypeHasModule() {
        var result = ModelUtils.GetModuleForObjectTypeRelation(GetTypeDefWithModules(), "relation4");
        Assert.Equal("type_module1", result);
    }

    [Fact]
    public void TestGetModuleForObjectTypeRelation_RelationDoesNotExist() {
        var exception = Assert.Throws<Exception>(() => {
            ModelUtils.GetModuleForObjectTypeRelation(GetTypeDefWithModules(), "relation5");
        });

        Assert.Equal("relation relation5 does not exist in type type1", exception.Message);
    }

    [Fact]
    public void TestGetModuleForObjectTypeRelation_RelationDoesNotExist2() {
        var exception = Assert.Throws<Exception>(() => {
            ModelUtils.GetModuleForObjectTypeRelation(GetTypeDefWithModules(), "relation6");
        });

        Assert.Equal("relation relation6 does not exist in type type1", exception.Message);
    }

    [Fact]
    public void TestGetModuleForObjectTypeRelation_RelationExistsButNoModule_TypeNoModule() {
        var result = ModelUtils.GetModuleForObjectTypeRelation(GetTypeDefWithoutModules(), "relation7");
        Assert.Null(result);
    }

    [Fact]
    public void TestIsRelationAssignable_RelationHasThis() {
        var relDef = new Userset {
            This = new object()
        };

        var result = ModelUtils.IsRelationAssignable(relDef);
        Assert.True(result);
    }

    [Fact]
    public void TestIsRelationAssignable_RelationHasUnionWithThis() {
        var relDef = new Userset {
            Union = new Usersets {
                Child = new List<Userset>
                {
                    new Userset
                    {
                        This = new object()
                    }
                }
            }
        };

        var result = ModelUtils.IsRelationAssignable(relDef);
        Assert.True(result);
    }

    [Fact]
    public void TestIsRelationAssignable_RelationHasIntersectionWithThis() {
        var relDef = new Userset {
            Intersection = new Usersets {
                Child = new List<Userset>
                {
                    new Userset
                    {
                        This = new object()
                    }
                }
            }
        };

        var result = ModelUtils.IsRelationAssignable(relDef);
        Assert.True(result);
    }

    [Fact]
    public void TestIsRelationAssignable_RelationHasDifferenceWithBaseThis() {
        var relDef = new Userset {
            Difference = new Difference {
                Base = new Userset {
                    This = new object()
                }
            }
        };

        var result = ModelUtils.IsRelationAssignable(relDef);
        Assert.True(result);
    }

    [Fact]
    public void TestIsRelationAssignable_RelationHasDifferenceWithSubtractThis() {
        var relDef = new Userset {
            Difference = new Difference {
                Subtract = new Userset {
                    This = new object()
                }
            }
        };

        var result = ModelUtils.IsRelationAssignable(relDef);
        Assert.True(result);
    }

    [Fact]
    public void TestIsRelationAssignable_RelationHasNoAssignableKeys() {
        var relDef = new Userset {
            Union = new Usersets {
                Child = new List<Userset>
                {
                    new Userset
                    {
                        Intersection = new Usersets
                        {
                            Child = new List<Userset>
                            {
                                new Userset()
                            }
                        }
                    }
                }
            }
        };

        var result = ModelUtils.IsRelationAssignable(relDef);
        Assert.False(result);
    }
}