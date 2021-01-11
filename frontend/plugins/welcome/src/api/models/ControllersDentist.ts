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
 * @interface ControllersDentist
 */
export interface ControllersDentist {
    /**
     * 
     * @type {string}
     * @memberof ControllersDentist
     */
    birthday?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentist
     */
    degree?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentist
     */
    expert?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentist
     */
    gender?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDentist
     */
    nurse?: number;
}

export function ControllersDentistFromJSON(json: any): ControllersDentist {
    return ControllersDentistFromJSONTyped(json, false);
}

export function ControllersDentistFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersDentist {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'birthday': !exists(json, 'birthday') ? undefined : json['birthday'],
        'degree': !exists(json, 'degree') ? undefined : json['degree'],
        'expert': !exists(json, 'expert') ? undefined : json['expert'],
        'gender': !exists(json, 'gender') ? undefined : json['gender'],
        'nurse': !exists(json, 'nurse') ? undefined : json['nurse'],
    };
}

export function ControllersDentistToJSON(value?: ControllersDentist | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'birthday': value.birthday,
        'degree': value.degree,
        'expert': value.expert,
        'gender': value.gender,
        'nurse': value.nurse,
    };
}

