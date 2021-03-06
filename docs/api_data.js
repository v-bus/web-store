define({ "api": [
  {
    "type": "DELETE",
    "url": "/",
    "title": "DELETE",
    "description": "<p>To destroy a resource and remove it from your account and environment, the DELETE method should be used. This will remove the specified object if it is found. If it is not found, the operation will return a response indicating that the object was not found. This idempotency means that you do not have to check for a resource's availability prior to issuing a delete command, the final state will be the same regardless of its existence.</p>",
    "version": "0.0.0",
    "group": "GlobalDescription",
    "filename": "./WebStore.js",
    "groupTitle": "GlobalDescription",
    "name": "Delete"
  },
  {
    "type": "GET",
    "url": "/",
    "title": "GET",
    "description": "<p>For simple retrieval of information about your account, SCO, SCO operations you should use the GET method. The information you request will be returned to you as a JSON object. The attributes defined by the JSON object can be used to form additional requests. Any request using the GET method is read-only and will not affect any of the objects you are querying.</p>",
    "version": "0.0.0",
    "group": "GlobalDescription",
    "filename": "./WebStore.js",
    "groupTitle": "GlobalDescription",
    "name": "Get"
  },
  {
    "type": "HTTP Statuses",
    "url": "/",
    "title": "HTTP Statuses",
    "description": "<p>Along with the HTTP methods that the API responds to, it will also return standard HTTP statuses, including error codes. In the event of a problem, the status will contain the error code, while the body of the response will usually contain additional information about the problem that was encountered. In general, if the status returned is in the 200 range, it indicates that the request was fulfilled successfully and that no error was encountered. Return codes in the 400 range typically indicate that there was an issue with the request that was sent. Among other things, this could mean that you did not authenticate correctly, that you are requesting an action that you do not have authorization for, that the object you are requesting does not exist, or that your request is malformed. If you receive a status in the 500 range, this generally indicates a server-side problem. This means that we are having an issue on our end and cannot fulfill your request currently.</p>",
    "version": "0.0.0",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "    HTTP/1.1 200 OK\n    Content-Type: application/json; charset=UTF-8\n    Date: Mon, 11 May 2020 15:11:20 GMT\nContent-Length: 918",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 403 Unauthorized\n Content-Type: application/json; charset=UTF-8\n Date: Mon, 11 May 2020 15:11:20 GMT",
          "type": "json"
        }
      ]
    },
    "group": "GlobalDescription",
    "filename": "./WebStore.js",
    "groupTitle": "GlobalDescription",
    "name": "Http statuses"
  },
  {
    "type": "POST",
    "url": "/",
    "title": "POST",
    "version": "0.0.0",
    "group": "GlobalDescription",
    "description": "<p>Create resource (product, users, consumers etc.)</p>",
    "filename": "./WebStore.js",
    "groupTitle": "GlobalDescription",
    "name": "Post"
  },
  {
    "type": "PUT",
    "url": "/",
    "title": "PUT",
    "version": "0.0.0",
    "group": "GlobalDescription",
    "description": "<p>Update resource (product, users, consumers etc.)</p>",
    "filename": "./WebStore.js",
    "groupTitle": "GlobalDescription",
    "name": "Put"
  },
  {
    "type": "Responses",
    "url": "/",
    "title": "Responses",
    "group": "GlobalDescription",
    "version": "0.0.0",
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "json",
            "optional": false,
            "field": "Responses",
            "description": "<p>Returned values or objects</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Single_Object:",
          "content": "{\n  \"id\"           : \"<UUID>\"\n  \"consumer_id\"  : \"500\"\n  \"firstname\"    : \"John\",\n  \"lastname\"     : \"Doe\"\n}",
          "type": "json"
        },
        {
          "title": "Object_Collection:",
          "content": "[  {\n \"id\": \"f2e19ada-9392-11ea-98dd-0242ac110002\",\n \"url\": \"http://shop.com\",\n \"title\": \"Title product\",\n \"price\": \"0.00\",\n \"currency\": \"TRY\",\n \"img_url\": \"http://shop.com\",\n \"created_at\": \"2020-05-11T14:23:13Z\",\n  \"last_track_at\": \"2020-05-11T14:23:13Z\"\n },\n {\n \"id\": \"f332e9ed-9392-11ea-98dd-0242ac110002\",\n \"url\": \"http://shop.com\",\n \"title\": \"Title product\",\n \"price\": \"0.00\",\n \"currency\": \"TRY\",\n \"img_url\": \"http://shop.com\",\n \"created_at\": \"2020-05-11T14:23:14Z\",\n \"last_track_at\": \"2020-05-11T14:23:14Z\"\n}]",
          "type": "json"
        }
      ]
    },
    "description": "<p>When a request is successful, a response body will typically be sent back in the form of a JSON object. An exception to this is when a DELETE request is processed, which will result in a successful HTTP 204 status and an empty response body. Inside of this JSON object, the resource root that was the target of the request will be set as the key. This will be the singular form of the word if the request operated on a single object, and the plural form of the word if a collection was processed. For example, if you send a GET request to /consumers/$consumer_ID you will get back an object with a key called &quot;consumer&quot;. However, if you send the GET request to the general collection at /consumers, you will get back an object with a key called &quot;consumers&quot;. The value of these keys will generally be a JSON object for a request on a single object and an array of objects for a request on a collection of objects.</p>",
    "filename": "./WebStore.js",
    "groupTitle": "GlobalDescription",
    "name": "Responses"
  },
  {
    "type": "POST",
    "url": "/product",
    "title": "Create product",
    "name": "create_product",
    "group": "Products",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>description</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "url",
            "description": "<p>URL to product page</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>Product title</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "price",
            "description": "<p>Product price</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "currency",
            "description": "<p>Price currency (is global for current user seccion)</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "img_url",
            "description": "<p>URL to image of the product</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"url\":\"http://shop.com/ipad\",\n\t \"title\":\"IPad\",\n\t \"price\":\"12.00\",\n\t \"currency\":\"RUB\",\n\t \"img_url\":\"http://shop.com/images/ipad.jpg\"\n\t}",
          "type": "type"
        }
      ]
    },
    "success": {
      "fields": {
        "Product information": [
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "id",
            "description": "<p>identifier of product</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "url",
            "description": "<p>URL to product page</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>Product title</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "price",
            "description": "<p>Product price</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "currency",
            "description": "<p>Price currency (is global for current user seccion)</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "img_url",
            "description": "<p>URL to image of the product</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "created_at",
            "description": "<p>date and time of product SCU record creation Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "last_buy_at",
            "description": "<p>date and time of the last order with this product Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t \"id\":\"0ca17c65-98f8-11ea-9b20-04d4c44d69ba\",\n\t \"url\":\"\",\n\t \"title\":\"\",\n\t \"price\":\"0.00\",\n\t \"currency\":\"TRY\",\n\t \"img_url\":\"\",\n\t \"created_at\":\"2020-05-18T14:09:31+03:00\",\n\t \"last_track_at\":\"2020-05-18T14:09:31+03:00\"\n\t}",
          "type": "json"
        }
      ]
    },
    "examples": [
      {
        "title": "cURL Example  usage:",
        "content": "curl -X POST -H \"Content-Type: application/json\" -H \"Authorization: Bearer $TOKEN\" -d '{\"url\":\"http://shop.com/ipad\",\"title\":\"IPad\",\"price\":\"12.00\",\"currency\":\"RUB\",\"img_url\":\"http://shop.com/images/ipad.jpg\"}' \"http://api.example.com/product\"",
        "type": "curl"
      }
    ],
    "description": "<p>Create new product</p>",
    "filename": "./http_crud_product.go",
    "groupTitle": "Products",
    "header": {
      "fields": {
        "Request Header Desciption": [
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Content-Type",
            "description": "<p>Specify content type</p>"
          },
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Authorization Token <code>Bearer</code> for users and additionally <code>Basic</code> for administrator</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Headers:",
          "content": "Content-Type: application/json\nAuthorization: Bearer <TOKEN DIGITS>",
          "type": "http"
        },
        {
          "title": "Response Headers:",
          "content": "HTTP/1.1 200 OK\nContent-Type: application/json; charset=UTF-8\nDate: Mon, 11 May 2020 15:11:20 GMT\nContent-Length: 918",
          "type": "HTTP"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 403 Unauthorized\n Content-Type: application/json; charset=UTF-8\n Date: Mon, 11 May 2020 15:11:20 GMT",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "DELETE",
    "url": "/product/:id",
    "title": "Delete product",
    "name": "delete_product",
    "group": "Products",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "id",
            "description": "<p>Product ID</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Status": [
          {
            "group": "Status",
            "type": "string",
            "optional": false,
            "field": "id",
            "description": "<p>Product ID</p>"
          },
          {
            "group": "Status",
            "type": "string",
            "optional": false,
            "field": "status",
            "description": "<p>Delete status should be &quot;deleted&quot;</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t \"id\":\"0ca17c65-98f8-11ea-9b20-04d4c44d69ba\",\n\t \"status\":\"deleted\"\n\t}",
          "type": "type"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response: No product in store",
          "content": "{\n\t \"id\":\"f31fd9f3-98ce-11ea-ab73-0242ac110002\",\n\t \"status\":\"error\",\n\t \"reason\":\"No such Product ID\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 403 Unauthorized\n Content-Type: application/json; charset=UTF-8\n Date: Mon, 11 May 2020 15:11:20 GMT",
          "type": "json"
        }
      ]
    },
    "examples": [
      {
        "title": "cURL Example  usage:",
        "content": "curl -v -X DELETE -H \"Authorization: Bearer 1234567890\" -H \"Content-Type: application/json\" \"http://api.example.com/product/edf2f8a2-9392-11ea-98dd-0242ac110002\"",
        "type": "curl"
      }
    ],
    "description": "<p>Update Product information. PUT <code>id</code> and new information structure.  Admin user only can update information.</p>",
    "filename": "./http_crud_product.go",
    "groupTitle": "Products",
    "header": {
      "fields": {
        "Request Header Desciption": [
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Content-Type",
            "description": "<p>Specify content type</p>"
          },
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Authorization Token <code>Bearer</code> for users and additionally <code>Basic</code> for administrator</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Headers:",
          "content": "Content-Type: application/json\nAuthorization: Bearer <TOKEN DIGITS>",
          "type": "http"
        },
        {
          "title": "Response Headers:",
          "content": "HTTP/1.1 200 OK\nContent-Type: application/json; charset=UTF-8\nDate: Mon, 11 May 2020 15:11:20 GMT\nContent-Length: 918",
          "type": "HTTP"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/product/:id",
    "title": "Product information",
    "name": "product_item",
    "group": "Products",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "id",
            "optional": false,
            "field": "ID",
            "description": "<p>unique ID of an item in store</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Product information": [
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "id",
            "description": "<p>identifier of product</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "url",
            "description": "<p>URL to product page</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>Product title</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "price",
            "description": "<p>Product price</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "currency",
            "description": "<p>Price currency (is global for current user seccion)</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "img_url",
            "description": "<p>URL to image of the product</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "created_at",
            "description": "<p>date and time of product SCU record creation Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "last_buy_at",
            "description": "<p>date and time of the last order with this product Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"id\": \"f332e9ed-9392-11ea-98dd-0242ac110002\",\n  \"url\": \"http://shop.com\",\n  \"title\": \"Title product\",\n  \"price\": \"12.00\",\n  \"currency\": \"RUB\",\n  \"img_url\": \"http://shop.com\",\n  \"created_at\": \"2020-05-11T14:23:14Z\",\n  \"last_track_at\": \"2020-05-11T14:23:14Z\"\n}",
          "type": "json"
        }
      ]
    },
    "examples": [
      {
        "title": "cURL Example  usage:",
        "content": "curl -X GET -H \"Content-Type: application/json\" -H \"Authorization: Bearer $TOKEN\" \"http://api.example.com/product/f332e9ed-9392-11ea-98dd-0242ac110002\"",
        "type": "curl"
      }
    ],
    "description": "<p>Detailed Product Information</p>",
    "filename": "./http_crud_product.go",
    "groupTitle": "Products",
    "header": {
      "fields": {
        "Request Header Desciption": [
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Content-Type",
            "description": "<p>Specify content type</p>"
          },
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Authorization Token <code>Bearer</code> for users and additionally <code>Basic</code> for administrator</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Headers:",
          "content": "Content-Type: application/json\nAuthorization: Bearer <TOKEN DIGITS>",
          "type": "http"
        },
        {
          "title": "Response Headers:",
          "content": "HTTP/1.1 200 OK\nContent-Type: application/json; charset=UTF-8\nDate: Mon, 11 May 2020 15:11:20 GMT\nContent-Length: 918",
          "type": "HTTP"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 403 Unauthorized\n Content-Type: application/json; charset=UTF-8\n Date: Mon, 11 May 2020 15:11:20 GMT",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/product/all",
    "title": "Products list",
    "name": "product_list",
    "group": "Products",
    "version": "0.1.0",
    "success": {
      "fields": {
        "Product information": [
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "id",
            "description": "<p>identifier of product</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "url",
            "description": "<p>URL to product page</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>Product title</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "price",
            "description": "<p>Product price</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "currency",
            "description": "<p>Price currency (is global for current user seccion)</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "img_url",
            "description": "<p>URL to image of the product</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "created_at",
            "description": "<p>date and time of product SCU record creation Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ</p>"
          },
          {
            "group": "Product information",
            "type": "string",
            "optional": false,
            "field": "last_buy_at",
            "description": "<p>date and time of the last order with this product Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "[  {\n \"id\": \"f2e19ada-9392-11ea-98dd-0242ac110002\",\n \"url\": \"http://shop.com\",\n \"title\": \"Title product\",\n \"price\": \"0.00\",\n \"currency\": \"TRY\",\n \"img_url\": \"http://shop.com\",\n \"created_at\": \"2020-05-11T14:23:13Z\",\n  \"last_track_at\": \"2020-05-11T14:23:13Z\"\n },\n {\n \"id\": \"f332e9ed-9392-11ea-98dd-0242ac110002\",\n \"url\": \"http://shop.com\",\n \"title\": \"Title product\",\n \"price\": \"0.00\",\n \"currency\": \"TRY\",\n \"img_url\": \"http://shop.com\",\n \"created_at\": \"2020-05-11T14:23:14Z\",\n \"last_track_at\": \"2020-05-11T14:23:14Z\"\n }]",
          "type": "type"
        }
      ]
    },
    "examples": [
      {
        "title": "cURL Example  usage:",
        "content": "curl -X GET -H \"Content-Type: application/json\" -H \"Authorization: Bearer $TOKEN\" \"http://api.example.com/product/all\"",
        "type": "curl"
      }
    ],
    "description": "<p>Product list</p>",
    "filename": "./http_crud_product.go",
    "groupTitle": "Products",
    "header": {
      "fields": {
        "Request Header Desciption": [
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Content-Type",
            "description": "<p>Specify content type</p>"
          },
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Authorization Token <code>Bearer</code> for users and additionally <code>Basic</code> for administrator</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Headers:",
          "content": "Content-Type: application/json\nAuthorization: Bearer <TOKEN DIGITS>",
          "type": "http"
        },
        {
          "title": "Response Headers:",
          "content": "HTTP/1.1 200 OK\nContent-Type: application/json; charset=UTF-8\nDate: Mon, 11 May 2020 15:11:20 GMT\nContent-Length: 918",
          "type": "HTTP"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 403 Unauthorized\n Content-Type: application/json; charset=UTF-8\n Date: Mon, 11 May 2020 15:11:20 GMT",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "PUT",
    "url": "/product/:id",
    "title": "Update Product Information",
    "name": "product_update",
    "group": "Products",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "url",
            "optional": false,
            "field": "url",
            "description": "<p>URL to product page</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>Product title</p>"
          },
          {
            "group": "Parameter",
            "type": "money",
            "optional": false,
            "field": "price",
            "description": "<p>Product price</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "currency",
            "description": "<p>Price currency (is global for current user seccion)</p>"
          },
          {
            "group": "Parameter",
            "type": "url",
            "optional": false,
            "field": "img_url",
            "description": "<p>URL to image of the product</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n \"id\": \"f332e9ed-9392-11ea-98dd-0242ac110002\",\n \"url\": \"http://shop.com\",\n \"title\": \"Title product\",\n \"price\": \"0.00\",\n \"currency\": \"TRY\",\n \"img_url\": \"http://shop.com\"\n}",
          "type": "type"
        }
      ]
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "longint",
            "optional": false,
            "field": "id",
            "description": "<p>Product ID</p>"
          },
          {
            "group": "200",
            "type": "string",
            "optional": false,
            "field": "status",
            "description": "<p>Result Responce <code>(success|fail|error|auth_error)</code>.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "\n{\"id\":\"edf2f8a2-9392-11ea-98dd-0242ac110002\",\"status\":\"success\"}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response: Failed means application errors",
          "content": "{\n     \"id\" : 34567876544\n     \"status\": \"fail\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response: Error means application errors",
          "content": "\n{\"id\":\"edf2f8a2-9392-11ea-98dd-0242ac11000\",\"status\":\"error\",\"reason\":\"No such Product ID\"}",
          "type": "json"
        },
        {
          "title": "Error-Response: No_User means no user was set in Header",
          "content": "{\n\t \"id\":\"f31fd9f3-98ce-11ea-ab73-0242ac110002\",\n  \"status\":\"error\",\n  \"reason\":\"No user specified\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response: Auth_Error means administrator authentication errors",
          "content": "{\n\t \"id\":\"f31fd9f3-98ce-11ea-ab73-0242ac110002\",\n\t \"status\":\"error\",\n\t \"reason\":\"Only admin can update product\"\n}",
          "type": "json"
        },
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 403 Unauthorized\n Content-Type: application/json; charset=UTF-8\n Date: Mon, 11 May 2020 15:11:20 GMT",
          "type": "json"
        }
      ]
    },
    "examples": [
      {
        "title": "cURL Example  usage:",
        "content": "curl -v -X PUT -H \"User: admin\" -H \"Authorization: Bearer 1234567890\" -H \"Content-Type: application/json\" -d '{\"url\":\"http://shop.com/my\",\"title\":\"Title product\",\"price\":\"0.00\",\"currency\":\"TRY\",\"img_url\":\"http://shop.com\"}' \"http://api.example.com/product/edf2f8a2-9392-11ea-98dd-0242ac110002\"",
        "type": "curl"
      }
    ],
    "description": "<p>Update Product information. PUT <code>id</code> and new information structure.  Admin user only can update information.</p>",
    "filename": "./http_crud_product.go",
    "groupTitle": "Products",
    "header": {
      "fields": {
        "Request Header Desciption": [
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Content-Type",
            "description": "<p>Specify content type</p>"
          },
          {
            "group": "Request Header Desciption",
            "type": "http",
            "optional": false,
            "field": "Authorization",
            "description": "<p>Authorization Token <code>Bearer</code> for users and additionally <code>Basic</code> for administrator</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Headers:",
          "content": "Content-Type: application/json\nAuthorization: Bearer <TOKEN DIGITS>",
          "type": "http"
        },
        {
          "title": "Response Headers:",
          "content": "HTTP/1.1 200 OK\nContent-Type: application/json; charset=UTF-8\nDate: Mon, 11 May 2020 15:11:20 GMT\nContent-Length: 918",
          "type": "HTTP"
        }
      ]
    }
  }
] });
