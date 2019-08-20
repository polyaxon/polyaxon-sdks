// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.7
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1EntityStatusRequest = factory(root.PolyaxonSdk.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The V1EntityStatusRequest model module.
   * @module model/V1EntityStatusRequest
   * @version 1.14.4
   */

  /**
   * Constructs a new <code>V1EntityStatusRequest</code>.
   * @alias module:model/V1EntityStatusRequest
   * @class
   */
  var exports = function() {
    var _this = this;





  };

  /**
   * Constructs a <code>V1EntityStatusRequest</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1EntityStatusRequest} obj Optional instance to populate.
   * @return {module:model/V1EntityStatusRequest} The populated <code>V1EntityStatusRequest</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('owner')) {
        obj['owner'] = ApiClient.convertToType(data['owner'], 'String');
      }
      if (data.hasOwnProperty('project')) {
        obj['project'] = ApiClient.convertToType(data['project'], 'String');
      }
      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('status')) {
        obj['status'] = ApiClient.convertToType(data['status'], 'String');
      }
    }
    return obj;
  }

  /**
   * @member {String} owner
   */
  exports.prototype['owner'] = undefined;
  /**
   * @member {String} project
   */
  exports.prototype['project'] = undefined;
  /**
   * @member {String} id
   */
  exports.prototype['id'] = undefined;
  /**
   * @member {String} status
   */
  exports.prototype['status'] = undefined;



  return exports;
}));

