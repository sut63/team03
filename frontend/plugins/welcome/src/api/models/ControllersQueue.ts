/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersQueue
 */
export interface ControllersQueue {
    /**
     * 
     * @type {string}
     * @memberof ControllersQueue
     */
    dental?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersQueue
     */
    dentist?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersQueue
     */
    patient?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersQueue
     */
    queueTime?: string;
}

export function ControllersQueueFromJSON(json: any): ControllersQueue {
    return ControllersQueueFromJSONTyped(json, false);
}

export function ControllersQueueFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersQueue {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dental': !exists(json, 'dental') ? undefined : json['dental'],
        'dentist': !exists(json, 'dentist') ? undefined : json['dentist'],
        'patient': !exists(json, 'patient') ? undefined : json['patient'],
        'queueTime': !exists(json, 'queueTime') ? undefined : json['queueTime'],
    };
}

export function ControllersQueueToJSON(value?: ControllersQueue | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dental': value.dental,
        'dentist': value.dentist,
        'patient': value.patient,
        'queueTime': value.queueTime,
    };
}


