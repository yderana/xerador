export default {
    post: {
        type: "object",
        properties: {
            name: {
                type: "string",
                pattern: "^[A-Za-z0-9 _)@./#&(+-]*$"
            },
        },
        required: ["name"],
        $schema: "http://json-schema.org/draft-07/schema#"
    },
    put: {
        type: "object",
        properties: {
            id: {
                type: "string",
                minLength: 1,
                pattern: "^[A-Za-z0-9 _)@./#&(+-]*$"
            },
            name: {
                type: "string",
                minLength: 1,
                pattern: "^[A-Za-z0-9 _)@./#&(+-]*$"
            }
        },
        required: ["id", "name"],
        $schema: "http://json-schema.org/draft-07/schema#"
    }
};
