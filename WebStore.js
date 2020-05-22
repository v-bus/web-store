/** 
 * @api {GET} / GET
 * @apiDescription For simple retrieval of information about your account, SCO, SCO operations you should use the GET method. The information you request will be returned to you as a JSON object.
 * The attributes defined by the JSON object can be used to form additional requests. Any request using the GET method is read-only and will not affect any of the objects you are querying.
 * @apiVersion  0.0.0
 * @apiGroup GlobalDescription
*/

/**
 * @api {DELETE} / DELETE
 * @apiDescription To destroy a resource and remove it from your account and environment, the DELETE method should be used. This will remove the specified object if it is found. If it is not found, the operation will return a response indicating that the object was not found.
 * This idempotency means that you do not have to check for a resource's availability prior to issuing a delete command, the final state will be the same regardless of its existence.
 * @apiVersion  0.0.0 
 * @apiGroup GlobalDescription
*/

/**
 * 
 * @api {PUT} / PUT
 * 
 * @apiVersion  0.0.0
 * @apiGroup GlobalDescription
 * 
 * @apiDescription Update resource (product, users, consumers etc.)
 */

/**
 * 
 * @api {POST} / POST
 * 
 * @apiVersion  0.0.0
 * @apiGroup GlobalDescription
 * 
 * @apiDescription Create resource (product, users, consumers etc.)
 */

/**
 * @api {HTTP Statuses} / HTTP Statuses
 * @apiDescription Along with the HTTP methods that the API responds to, it will also return standard HTTP statuses, including error codes.
 * In the event of a problem, the status will contain the error code, while the body of the response will usually contain additional information about the problem that was encountered.
 * In general, if the status returned is in the 200 range, it indicates that the request was fulfilled successfully and that no error was encountered.
 * Return codes in the 400 range typically indicate that there was an issue with the request that was sent. Among other things, this could mean that you did not authenticate correctly, that you are requesting an action that you do not have authorization for, that the object you are requesting does not exist, or that your request is malformed.
 * If you receive a status in the 500 range, this generally indicates a server-side problem. This means that we are having an issue on our end and cannot fulfill your request currently.
 * @apiVersion  0.0.0 
 * @apiSuccessExample {json} Success-Response:
 *     HTTP/1.1 200 OK
 *     Content-Type: application/json; charset=UTF-8
 *     Date: Mon, 11 May 2020 15:11:20 GMT
 * Content-Length: 918
 *  @apiErrorExample {json} Error-Response:
 *    HTTP/1.1 403 Unauthorized
 *     Content-Type: application/json; charset=UTF-8
 *     Date: Mon, 11 May 2020 15:11:20 GMT
 * @apiGroup GlobalDescription
*/

/**
 * @api {Responses} / Responses
 * @apiGroup GlobalDescription
 * @apiVersion  0.0.0
 * @apiSuccess (200) {json} Responses Returned values or objects
 * @apiSuccessExample {json} Single_Object:
 *     {
 *       "id"           : "<UUID>"
 *       "consumer_id"  : "500"
 *       "firstname"    : "John",
 *       "lastname"     : "Doe"
 *     }
 * @apiSuccessExample {json} Object_Collection:
 * [  {
 *  "id": "f2e19ada-9392-11ea-98dd-0242ac110002",
 *  "url": "http://shop.com",
 *  "title": "Title product",
 *  "price": "0.00",
 *  "currency": "TRY",
 *  "img_url": "http://shop.com",
 *  "created_at": "2020-05-11T14:23:13Z",
 *   "last_track_at": "2020-05-11T14:23:13Z"
 *  },
 *  {
 *  "id": "f332e9ed-9392-11ea-98dd-0242ac110002",
 *  "url": "http://shop.com",
 *  "title": "Title product",
 *  "price": "0.00",
 *  "currency": "TRY",
 *  "img_url": "http://shop.com",
 *  "created_at": "2020-05-11T14:23:14Z",
 *  "last_track_at": "2020-05-11T14:23:14Z"
 *}]
 *
 * @apiDescription When a request is successful, a response body will typically be sent back in the form of a JSON object. An exception to this is when a DELETE request is processed, which will result in a successful HTTP 204 status and an empty response body.
 * Inside of this JSON object, the resource root that was the target of the request will be set as the key. This will be the singular form of the word if the request operated on a single object, and the plural form of the word if a collection was processed.
 * For example, if you send a GET request to /consumers/$consumer_ID you will get back an object with a key called "consumer". However, if you send the GET request to the general collection at /consumers, you will get back an object with a key called "consumers".
 * The value of these keys will generally be a JSON object for a request on a single object and an array of objects for a request on a collection of objects.
 */

 /**
  * 
  * @apiDefine ReqHeaders used headers
  *    HTTP headers to make requests
  * @apiHeader (Request Header Desciption) {http} Content-Type Specify content type 
  * @apiHeader (Request Header Desciption) {http} Authorization  Authorization Token `Bearer` for users and additionally `Basic` for administrator
  * @apiHeaderExample {http} Request Headers:
  *    Content-Type: application/json
  *    Authorization: Bearer <TOKEN DIGITS>
  * @apiErrorExample {json} Error-Response:
  *    HTTP/1.1 403 Unauthorized
  *     Content-Type: application/json; charset=UTF-8
  *     Date: Mon, 11 May 2020 15:11:20 GMT
  * @apiHeaderExample {HTTP} Response Headers:
 *     HTTP/1.1 200 OK
 *     Content-Type: application/json; charset=UTF-8
 *     Date: Mon, 11 May 2020 15:11:20 GMT
 *     Content-Length: 918
  */