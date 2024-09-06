// ModelUtilsTest.java
package dev.openfga.language.utils;

import static org.junit.jupiter.api.Assertions.*;

import dev.openfga.sdk.api.model.*;
import java.util.List;
import java.util.Map;
import org.junit.jupiter.api.Test;

public class ModelUtilsTest {
    private TypeDefinition getTypeDefWithModules() {
        return new TypeDefinition()
                .type("type1")
                .relations(Map.of(
                        "relation1", new Userset(),
                        "relation2", new Userset(),
                        "relation3", new Userset(),
                        "relation4", new Userset()))
                .metadata(new Metadata()
                        .module("type_module1")
                        .relations(Map.of(
                                "relation1",
                                new RelationMetadata().module("module1"),
                                "relation2",
                                new RelationMetadata().module(""),
                                "relation3",
                                new RelationMetadata(),
                                "relation5",
                                new RelationMetadata())));
    }

    private TypeDefinition getTypeDefWithoutModules() {
        return new TypeDefinition().type("type2").relations(Map.of("relation7", new Userset()));
    }

    @Test
    public void testGetModuleForObjectTypeRelation_RelationExistsAndHasModule() throws Exception {
        String result = ModelUtils.getModuleForObjectTypeRelation(getTypeDefWithModules(), "relation1");
        assertEquals("module1", result);
    }

    @Test
    public void testGetModuleForObjectTypeRelation_RelationExistsButHasEmptyModule_TypeHasModule() throws Exception {
        String result = ModelUtils.getModuleForObjectTypeRelation(getTypeDefWithModules(), "relation2");
        assertEquals("type_module1", result);
    }

    @Test
    public void testGetModuleForObjectTypeRelation_RelationExistsButNoModule_TypeHasModule() throws Exception {
        String result = ModelUtils.getModuleForObjectTypeRelation(getTypeDefWithModules(), "relation3");
        assertEquals("type_module1", result);
    }

    @Test
    public void testGetModuleForObjectTypeRelation_RelationExistsButNoMetadata_TypeHasModule() throws Exception {
        String result = ModelUtils.getModuleForObjectTypeRelation(getTypeDefWithModules(), "relation4");
        assertEquals("type_module1", result);
    }

    @Test
    public void testGetModuleForObjectTypeRelation_RelationDoesNotExist() {
        Exception exception = assertThrows(Exception.class, () -> {
            ModelUtils.getModuleForObjectTypeRelation(getTypeDefWithModules(), "relation5");
        });

        assertEquals("relation relation5 does not exist in type type1", exception.getMessage());
    }

    @Test
    public void testGetModuleForObjectTypeRelation_RelationDoesNotExist2() {
        Exception exception = assertThrows(Exception.class, () -> {
            ModelUtils.getModuleForObjectTypeRelation(getTypeDefWithModules(), "relation6");
        });

        assertEquals("relation relation6 does not exist in type type1", exception.getMessage());
    }

    @Test
    public void testGetModuleForObjectTypeRelation_RelationExistsButNoModule_TypeNoModule() throws Exception {
        String result = ModelUtils.getModuleForObjectTypeRelation(getTypeDefWithoutModules(), "relation7");
        assertEquals(null, result);
    }

    @Test
    public void testIsRelationAssignable_RelationHasThis() {
        Userset relDef = new Userset();
        relDef.setThis(new Object());

        boolean result = ModelUtils.isRelationAssignable(relDef);
        assertTrue(result);
    }

    @Test
    public void testIsRelationAssignable_RelationHasUnionWithThis() {
        Userset relDef = new Userset();
        Usersets union = new Usersets();
        Userset innerRelDef = new Userset();
        innerRelDef.setThis(new Object());
        union.setChild(List.of(innerRelDef));
        relDef.setUnion(union);

        boolean result = ModelUtils.isRelationAssignable(relDef);
        assertTrue(result);
    }

    @Test
    public void testIsRelationAssignable_RelationHasIntersectionWithThis() {
        Userset relDef = new Userset();
        Usersets intersection = new Usersets();
        Userset innerRelDef = new Userset();
        innerRelDef.setThis(new Object());
        intersection.setChild(List.of(innerRelDef));
        relDef.setIntersection(intersection);

        boolean result = ModelUtils.isRelationAssignable(relDef);
        assertTrue(result);
    }

    @Test
    public void testIsRelationAssignable_RelationHasDifferenceWithBaseThis() {
        Userset relDef = new Userset();
        Difference difference = new Difference();
        Userset innerRelDef = new Userset();
        innerRelDef.setThis(new Object());
        difference.setBase(innerRelDef);
        relDef.setDifference(difference);

        boolean result = ModelUtils.isRelationAssignable(relDef);
        assertTrue(result);
    }

    @Test
    public void testIsRelationAssignable_RelationHasDifferenceWithSubtractThis() {
        Userset relDef = new Userset();
        Difference difference = new Difference();
        Userset innerRelDef = new Userset();
        innerRelDef.setThis(new Object());
        difference.setSubtract(innerRelDef);
        relDef.setDifference(difference);

        boolean result = ModelUtils.isRelationAssignable(relDef);
        assertTrue(result);
    }

    @Test
    public void testIsRelationAssignable_RelationHasNoAssignableKeys() {
        Userset relDef = new Userset();
        Usersets union = new Usersets();
        Usersets intersection = new Usersets();
        intersection.setChild(List.of(new Userset()));
        Userset intersectionUserset = new Userset();
        intersectionUserset.setIntersection(intersection);
        union.setChild(List.of(intersectionUserset));
        relDef.setUnion(union);

        boolean result = ModelUtils.isRelationAssignable(relDef);
        assertFalse(result);
    }
}
