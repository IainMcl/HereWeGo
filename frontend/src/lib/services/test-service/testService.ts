import { request } from "$lib/services/api-service/apiService";

export async function test(): Promise<string> {
    const response = await request('/api/test', 'GET');
    return response;
}