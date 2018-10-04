/**
 * Xegony API
 * generic description goes here
 *
 * OpenAPI spec version: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


import * as $ from 'jquery';
import * as models from '../model/models';
import { COLLECTION_FORMATS } from '../variables';
import { Configuration } from '../configuration';

/* tslint:disable:no-unused-variable member-ordering */


export class XegonyApi {
    protected basePath = 'http://localhost';
    public defaultHeaders: Array<string> = [];
    public defaultExtraJQueryAjaxSettings?: JQueryAjaxSettings = null;
    public configuration: Configuration = new Configuration();

    constructor(basePath?: string, configuration?: Configuration, defaultExtraJQueryAjaxSettings?: JQueryAjaxSettings) {
        if (basePath) {
            this.basePath = basePath;
        }
        if (configuration) {
            this.configuration = configuration;
        }
        if (defaultExtraJQueryAjaxSettings) {
            this.defaultExtraJQueryAjaxSettings = defaultExtraJQueryAjaxSettings;
        }
    }

    private extendObj<T1, T2 extends T1>(objA: T2, objB: T2): T1|T2 {
        for (let key in objB) {
            if (objB.hasOwnProperty(key)) {
                objA[key] = objB[key];
            }
        }
        return objA;
    }

    /**
     * 
     * @param body 
     */
    public npcCreate(body: models.PbNpcCreateRequest, extraJQueryAjaxSettings?: JQueryAjaxSettings): JQueryPromise<{ response: JQueryXHR; body: models.PbNpcCreateResponse;  }> {
        let localVarPath = this.basePath + '/v1/npc';

        let queryParameters: any = {};
        let headerParams: any = {};
        // verify required parameter 'body' is not null or undefined
        if (body === null || body === undefined) {
            throw new Error('Required parameter body was null or undefined when calling npcCreate.');
        }


        localVarPath = localVarPath + "?" + $.param(queryParameters);
        // to determine the Content-Type header
        let consumes: string[] = [
            'application/json'
        ];

        // to determine the Accept header
        let produces: string[] = [
            'application/json'
        ];


        headerParams['Content-Type'] = 'application/json';

        let requestOptions: JQueryAjaxSettings = {
            url: localVarPath,
            type: 'POST',
            headers: headerParams,
            processData: false
        };

        requestOptions.data = JSON.stringify(body);
        if (headerParams['Content-Type']) {
            requestOptions.contentType = headerParams['Content-Type'];
        }

        if (extraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, extraJQueryAjaxSettings);
        }

        if (this.defaultExtraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, this.defaultExtraJQueryAjaxSettings);
        }

        let dfd = $.Deferred();
        $.ajax(requestOptions).then(
            (data: models.PbNpcCreateResponse, textStatus: string, jqXHR: JQueryXHR) =>
                dfd.resolve(jqXHR, data),
            (xhr: JQueryXHR, textStatus: string, errorThrown: string) =>
                dfd.reject(xhr, errorThrown)
        );
        return dfd.promise();
    }

    /**
     * 
     * @param id 
     */
    public npcDelete(id: string, extraJQueryAjaxSettings?: JQueryAjaxSettings): JQueryPromise<{ response: JQueryXHR; body: models.PbNpcDeleteResponse;  }> {
        let localVarPath = this.basePath + '/v1/npc/{id}'.replace('{' + 'id' + '}', encodeURIComponent(String(id)));

        let queryParameters: any = {};
        let headerParams: any = {};
        // verify required parameter 'id' is not null or undefined
        if (id === null || id === undefined) {
            throw new Error('Required parameter id was null or undefined when calling npcDelete.');
        }


        localVarPath = localVarPath + "?" + $.param(queryParameters);
        // to determine the Content-Type header
        let consumes: string[] = [
            'application/json'
        ];

        // to determine the Accept header
        let produces: string[] = [
            'application/json'
        ];


        let requestOptions: JQueryAjaxSettings = {
            url: localVarPath,
            type: 'DELETE',
            headers: headerParams,
            processData: false
        };

        if (headerParams['Content-Type']) {
            requestOptions.contentType = headerParams['Content-Type'];
        }

        if (extraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, extraJQueryAjaxSettings);
        }

        if (this.defaultExtraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, this.defaultExtraJQueryAjaxSettings);
        }

        let dfd = $.Deferred();
        $.ajax(requestOptions).then(
            (data: models.PbNpcDeleteResponse, textStatus: string, jqXHR: JQueryXHR) =>
                dfd.resolve(jqXHR, data),
            (xhr: JQueryXHR, textStatus: string, errorThrown: string) =>
                dfd.reject(xhr, errorThrown)
        );
        return dfd.promise();
    }

    /**
     * 
     * @param id 
     * @param body 
     */
    public npcPatch(id: string, body: models.PbNpcPatchRequest, extraJQueryAjaxSettings?: JQueryAjaxSettings): JQueryPromise<{ response: JQueryXHR; body: models.PbNpcPatchResponse;  }> {
        let localVarPath = this.basePath + '/v1/npc/{id}'.replace('{' + 'id' + '}', encodeURIComponent(String(id)));

        let queryParameters: any = {};
        let headerParams: any = {};
        // verify required parameter 'id' is not null or undefined
        if (id === null || id === undefined) {
            throw new Error('Required parameter id was null or undefined when calling npcPatch.');
        }

        // verify required parameter 'body' is not null or undefined
        if (body === null || body === undefined) {
            throw new Error('Required parameter body was null or undefined when calling npcPatch.');
        }


        localVarPath = localVarPath + "?" + $.param(queryParameters);
        // to determine the Content-Type header
        let consumes: string[] = [
            'application/json'
        ];

        // to determine the Accept header
        let produces: string[] = [
            'application/json'
        ];


        headerParams['Content-Type'] = 'application/json';

        let requestOptions: JQueryAjaxSettings = {
            url: localVarPath,
            type: 'PATCH',
            headers: headerParams,
            processData: false
        };

        requestOptions.data = JSON.stringify(body);
        if (headerParams['Content-Type']) {
            requestOptions.contentType = headerParams['Content-Type'];
        }

        if (extraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, extraJQueryAjaxSettings);
        }

        if (this.defaultExtraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, this.defaultExtraJQueryAjaxSettings);
        }

        let dfd = $.Deferred();
        $.ajax(requestOptions).then(
            (data: models.PbNpcPatchResponse, textStatus: string, jqXHR: JQueryXHR) =>
                dfd.resolve(jqXHR, data),
            (xhr: JQueryXHR, textStatus: string, errorThrown: string) =>
                dfd.reject(xhr, errorThrown)
        );
        return dfd.promise();
    }

    /**
     * 
     * @param id 
     */
    public npcRead(id: string, extraJQueryAjaxSettings?: JQueryAjaxSettings): JQueryPromise<{ response: JQueryXHR; body: models.PbNpcReadResponse;  }> {
        let localVarPath = this.basePath + '/v1/npc/{id}'.replace('{' + 'id' + '}', encodeURIComponent(String(id)));

        let queryParameters: any = {};
        let headerParams: any = {};
        // verify required parameter 'id' is not null or undefined
        if (id === null || id === undefined) {
            throw new Error('Required parameter id was null or undefined when calling npcRead.');
        }


        localVarPath = localVarPath + "?" + $.param(queryParameters);
        // to determine the Content-Type header
        let consumes: string[] = [
            'application/json'
        ];

        // to determine the Accept header
        let produces: string[] = [
            'application/json'
        ];


        let requestOptions: JQueryAjaxSettings = {
            url: localVarPath,
            type: 'GET',
            headers: headerParams,
            processData: false
        };

        if (headerParams['Content-Type']) {
            requestOptions.contentType = headerParams['Content-Type'];
        }

        if (extraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, extraJQueryAjaxSettings);
        }

        if (this.defaultExtraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, this.defaultExtraJQueryAjaxSettings);
        }

        let dfd = $.Deferred();
        $.ajax(requestOptions).then(
            (data: models.PbNpcReadResponse, textStatus: string, jqXHR: JQueryXHR) =>
                dfd.resolve(jqXHR, data),
            (xhr: JQueryXHR, textStatus: string, errorThrown: string) =>
                dfd.reject(xhr, errorThrown)
        );
        return dfd.promise();
    }

    /**
     * 
     * @param name 
     * @param limit 
     * @param offset 
     */
    public npcSearch(name: string, limit?: string, offset?: string, extraJQueryAjaxSettings?: JQueryAjaxSettings): JQueryPromise<{ response: JQueryXHR; body: models.PbNpcSearchResponse;  }> {
        let localVarPath = this.basePath + '/v1/npc/search/{name}'.replace('{' + 'name' + '}', encodeURIComponent(String(name)));

        let queryParameters: any = {};
        let headerParams: any = {};
        // verify required parameter 'name' is not null or undefined
        if (name === null || name === undefined) {
            throw new Error('Required parameter name was null or undefined when calling npcSearch.');
        }

        if (limit !== null && limit !== undefined) {
            queryParameters['limit'] = <string><any>limit;
        }
        if (offset !== null && offset !== undefined) {
            queryParameters['offset'] = <string><any>offset;
        }

        localVarPath = localVarPath + "?" + $.param(queryParameters);
        // to determine the Content-Type header
        let consumes: string[] = [
            'application/json'
        ];

        // to determine the Accept header
        let produces: string[] = [
            'application/json'
        ];


        let requestOptions: JQueryAjaxSettings = {
            url: localVarPath,
            type: 'GET',
            headers: headerParams,
            processData: false
        };

        if (headerParams['Content-Type']) {
            requestOptions.contentType = headerParams['Content-Type'];
        }

        if (extraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, extraJQueryAjaxSettings);
        }

        if (this.defaultExtraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, this.defaultExtraJQueryAjaxSettings);
        }

        let dfd = $.Deferred();
        $.ajax(requestOptions).then(
            (data: models.PbNpcSearchResponse, textStatus: string, jqXHR: JQueryXHR) =>
                dfd.resolve(jqXHR, data),
            (xhr: JQueryXHR, textStatus: string, errorThrown: string) =>
                dfd.reject(xhr, errorThrown)
        );
        return dfd.promise();
    }

    /**
     * 
     * @param id 
     * @param body 
     */
    public npcUpdate(id: string, body: models.PbNpcUpdateRequest, extraJQueryAjaxSettings?: JQueryAjaxSettings): JQueryPromise<{ response: JQueryXHR; body: models.PbNpcUpdateResponse;  }> {
        let localVarPath = this.basePath + '/v1/npc/{id}'.replace('{' + 'id' + '}', encodeURIComponent(String(id)));

        let queryParameters: any = {};
        let headerParams: any = {};
        // verify required parameter 'id' is not null or undefined
        if (id === null || id === undefined) {
            throw new Error('Required parameter id was null or undefined when calling npcUpdate.');
        }

        // verify required parameter 'body' is not null or undefined
        if (body === null || body === undefined) {
            throw new Error('Required parameter body was null or undefined when calling npcUpdate.');
        }


        localVarPath = localVarPath + "?" + $.param(queryParameters);
        // to determine the Content-Type header
        let consumes: string[] = [
            'application/json'
        ];

        // to determine the Accept header
        let produces: string[] = [
            'application/json'
        ];


        headerParams['Content-Type'] = 'application/json';

        let requestOptions: JQueryAjaxSettings = {
            url: localVarPath,
            type: 'PUT',
            headers: headerParams,
            processData: false
        };

        requestOptions.data = JSON.stringify(body);
        if (headerParams['Content-Type']) {
            requestOptions.contentType = headerParams['Content-Type'];
        }

        if (extraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, extraJQueryAjaxSettings);
        }

        if (this.defaultExtraJQueryAjaxSettings) {
            requestOptions = (<any>Object).assign(requestOptions, this.defaultExtraJQueryAjaxSettings);
        }

        let dfd = $.Deferred();
        $.ajax(requestOptions).then(
            (data: models.PbNpcUpdateResponse, textStatus: string, jqXHR: JQueryXHR) =>
                dfd.resolve(jqXHR, data),
            (xhr: JQueryXHR, textStatus: string, errorThrown: string) =>
                dfd.reject(xhr, errorThrown)
        );
        return dfd.promise();
    }

}