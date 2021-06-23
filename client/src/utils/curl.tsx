import React from 'react';
import axios from 'axios';

export const Curl = async (url: string, data:any, headers:any) =>{
    return await axios.post(url, data, {
        headers:headers
    });
}

export const getCurl = async (url: string, headers:any) =>{
    return await axios.get(url, {
        headers:headers
    });
}