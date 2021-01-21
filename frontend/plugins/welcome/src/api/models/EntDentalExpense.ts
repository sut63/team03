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
import {
    EntDentalexpenseEdges,
    EntDentalexpenseEdgesFromJSON,
    EntDentalexpenseEdgesFromJSONTyped,
    EntDentalexpenseEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDentalexpense
 */
export interface EntDentalexpense {
    /**
     * AddedTime holds the value of the "AddedTime" field.
     * @type {string}
     * @memberof EntDentalexpense
     */
    addedTime?: string;
    /**
     * Name holds the value of the "Name" field.
     * @type {string}
     * @memberof EntDentalexpense
     */
    name?: string;
    /**
     * Phone holds the value of the "Phone" field.
     * @type {string}
     * @memberof EntDentalexpense
     */
    phone?: string;
    /**
     * Rates holds the value of the "Rates" field.
     * @type {number}
     * @memberof EntDentalexpense
     */
    rates?: number;
    /**
     * Tax holds the value of the "Tax" field.
     * @type {string}
     * @memberof EntDentalexpense
     */
    tax?: string;
    /**
     * 
     * @type {EntDentalexpenseEdges}
     * @memberof EntDentalexpense
     */
    edges?: EntDentalexpenseEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDentalexpense
     */
    id?: number;
}

export function EntDentalexpenseFromJSON(json: any): EntDentalexpense {
    return EntDentalexpenseFromJSONTyped(json, false);
}

export function EntDentalexpenseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDentalexpense {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'AddedTime') ? undefined : json['AddedTime'],
        'name': !exists(json, 'Name') ? undefined : json['Name'],
        'phone': !exists(json, 'Phone') ? undefined : json['Phone'],
        'rates': !exists(json, 'Rates') ? undefined : json['Rates'],
        'tax': !exists(json, 'Tax') ? undefined : json['Tax'],
        'edges': !exists(json, 'edges') ? undefined : EntDentalexpenseEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDentalexpenseToJSON(value?: EntDentalexpense | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'AddedTime': value.addedTime,
        'Name': value.name,
        'Phone': value.phone,
        'Rates': value.rates,
        'Tax': value.tax,
        'edges': EntDentalexpenseEdgesToJSON(value.edges),
        'id': value.id,
    };
}


