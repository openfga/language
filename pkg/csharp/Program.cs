using OpenFgaLanguage.Transformers;
using System.Text.Json;

var dsl = """
          model
            schema 1.1

          type organization
            relations
              define member: [user] or owner
              define owner: [user]
              define repo_admin: [user, organization#member]
              define repo_reader: [user, organization#member]
              define repo_writer: [user, organization#member]

          type repo
            relations
              define admin: [user, team#member] or repo_admin from owner
              define maintainer: [user, team#member] or admin
              define owner: [organization]
              define reader: [user, team#member] or triager or repo_reader from owner
              define triager: [user, team#member] or writer
              define writer: [user, team#member] or maintainer or repo_writer from owner

          type team
            relations
              define member: [user, team#member]

          type user

          """;

var result = new DslToJsonTransformer().ParseDsl(dsl);

var json = JsonSerializer.Serialize(result.AuthorizationModel, new JsonSerializerOptions()
{
    WriteIndented = true
});


return;




/*
{
     "schema_version": "1.1",
     "type_definitions": [
       {
         "type": "organization",
         "relations": {
           "member": {
             "union": {
               "child": [
                 {
                   "this": {}
                 },
                 {
                   "computedUserset": {
                     "object": "",
                     "relation": "owner"
                   }
                 }
               ]
             }
           },
           "owner": {
             "this": {}
           },
           "repo_admin": {
             "this": {}
           },
           "repo_reader": {
             "this": {}
           },
           "repo_writer": {
             "this": {}
           }
         },
         "metadata": {
           "relations": {
             "member": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "owner": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "repo_admin": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "organization",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "repo_reader": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "organization",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "repo_writer": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "organization",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             }
           },
           "module": "",
           "source_info": null
         }
       },
       {
         "type": "repo",
         "relations": {
           "admin": {
             "union": {
               "child": [
                 {
                   "this": {}
                 },
                 {
                   "tupleToUserset": {
                     "tupleset": {
                       "object": "",
                       "relation": "owner"
                     },
                     "computedUserset": {
                       "object": "",
                       "relation": "repo_admin"
                     }
                   }
                 }
               ]
             }
           },
           "maintainer": {
             "union": {
               "child": [
                 {
                   "this": {}
                 },
                 {
                   "computedUserset": {
                     "object": "",
                     "relation": "admin"
                   }
                 }
               ]
             }
           },
           "owner": {
             "this": {}
           },
           "reader": {
             "union": {
               "child": [
                 {
                   "this": {}
                 },
                 {
                   "computedUserset": {
                     "object": "",
                     "relation": "triager"
                   }
                 },
                 {
                   "tupleToUserset": {
                     "tupleset": {
                       "object": "",
                       "relation": "owner"
                     },
                     "computedUserset": {
                       "object": "",
                       "relation": "repo_reader"
                     }
                   }
                 }
               ]
             }
           },
           "triager": {
             "union": {
               "child": [
                 {
                   "this": {}
                 },
                 {
                   "computedUserset": {
                     "object": "",
                     "relation": "writer"
                   }
                 }
               ]
             }
           },
           "writer": {
             "union": {
               "child": [
                 {
                   "this": {}
                 },
                 {
                   "computedUserset": {
                     "object": "",
                     "relation": "maintainer"
                   }
                 },
                 {
                   "tupleToUserset": {
                     "tupleset": {
                       "object": "",
                       "relation": "owner"
                     },
                     "computedUserset": {
                       "object": "",
                       "relation": "repo_writer"
                     }
                   }
                 }
               ]
             }
           }
         },
         "metadata": {
           "relations": {
             "admin": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "team",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "maintainer": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "team",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "owner": {
               "directly_related_user_types": [
                 {
                   "type": "organization",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "reader": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "team",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "triager": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "team",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             },
             "writer": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "team",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             }
           },
           "module": "",
           "source_info": null
         }
       },
       {
         "type": "team",
         "relations": {
           "member": {
             "this": {}
           }
         },
         "metadata": {
           "relations": {
             "member": {
               "directly_related_user_types": [
                 {
                   "type": "user",
                   "condition": ""
                 },
                 {
                   "type": "team",
                   "relation": "member",
                   "condition": ""
                 }
               ],
               "module": "",
               "source_info": null
             }
           },
           "module": "",
           "source_info": null
         }
       },
       {
         "type": "user",
         "relations": {},
         "metadata": null
       }
     ],
     "conditions": {}
   }
 */
