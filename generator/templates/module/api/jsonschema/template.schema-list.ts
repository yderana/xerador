export default {
    graphql: {
        type: "object",
        properties: {
            page: {
                type: "number"
            },
            limit: {
                type: "number"
            },
            orderBy: {
                type: "string",
                pattern: "^(id|name|createdAt|modifiedAt)$"
            },
            sort: {
                type: "string",
                pattern: "^(ASC|DESC|asc|desc)$"
            },
            noCache: {
                type: "string",
                pattern: "^(True|False|TRUE|FALSE|true|false)$"
            },
            query: {
                type: "string",
                pattern: "^[ A-Za-z0-9_)@./#&(+-]*$"
            }
        },
        required: ["page", "limit"],
        $schema: "http://json-schema.org/draft-07/schema#"
    },
    rest: {
        type: "object",
        properties: {
            page: {
                type: "string",
                pattern: "^[0-9]*$"
            },
            limit: {
                type: "string",
                pattern: "^[0-9]*$"
            },
            orderBy: {
                type: "string",
                pattern: "^(id|name|createdAt|modifiedAt)$"
            },
            sort: {
                type: "string",
                pattern: "^(ASC|DESC|asc|desc)$"
            },
            noCache: {
                type: "string",
                pattern: "^(True|False|TRUE|FALSE|true|false)$"
            },
            query: {
                type: "string",
                pattern: "^[ A-Za-z0-9_)@./#&(+-]*$"
            }
        },
        required: ["page", "limit"],
        $schema: "http://json-schema.org/draft-07/schema#"
    }
};
