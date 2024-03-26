import { request } from '$lib/services/api-service/apiService';

export const load = async (event) => {
    const data = await request("GET", "/health", null);
    return {
        data,
        status: 200,
    };
}