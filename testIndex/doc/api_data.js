define({ "api": [
  {
    "type": "GET",
    "url": "/product/:type",
    "title": "Get All Documents from TEST INDEX DOC",
    "name": "GetAllType",
    "group": "Product",
    "version": "0.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "paramName",
            "description": "<p>description</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    property : value\n}",
          "type": "type"
        }
      ]
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "type",
            "optional": false,
            "field": "name",
            "description": "<p>description</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    property : value\n}",
          "type": "type"
        }
      ]
    },
    "filename": "testIndex/trash.go",
    "groupTitle": "Product"
  }
] });
