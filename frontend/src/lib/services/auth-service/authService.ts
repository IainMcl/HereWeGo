import { request, unauthenticatedRequest } from '$lib/services/api-service/apiService';



/**
 * Logs out a user and removes the token from cookies
 * @returns 
 */
export async function logout(): Promise<any> {
    return { message: 'Logged out' };
}

/**
 * Logs in a user and stores the token in a cookie
 * @param email 
 * @param password 
 * @returns 
 */
export async function login(email: string, password: string): Promise<any> {
    const resp = await unauthenticatedRequest('POST', '/auth/login', { email, password });
    return resp;
}