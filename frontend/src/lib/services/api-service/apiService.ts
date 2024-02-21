/**
 * Makes a request to the API endpoint 
 * @param url Relative path to the API endpoint e.g. /api/notes
 * @param method HTTP method e.g. GET, POST, PUT, DELETE
 * @param data Data to be sent to the API endpoint
 * @returns JSON response from the API endpoint
 */
export async function request(url: string, method: string = 'GET', data: any = null): Promise<any> {
    const apiRoot = import.meta.env.VITE_API_ROOT;
    const fullUrl = apiRoot + url;
    const options: RequestInit = {
        method,
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token') || ''
        }
    };

    if (data) {
        options.body = JSON.stringify(data);
    }

    const response = await fetch(fullUrl, options);
    if (!response.ok) {
        throw new Error(`HTTP error: status: ${response.status}`);
    } else if (response.status === 204) {
        return;
    }
    return await response.json();
}

export async function unauthenticatedRequest(url: string, method: string = 'GET', data: any = null): Promise<any> {
    const apiRoot = import.meta.env.VITE_API_ROOT;
    const fullUrl = apiRoot + url;
    const options: RequestInit = {
        method,
        headers: {
            'Content-Type': 'application/json',
        }
    };

    if (data) {
        options.body = JSON.stringify(data);
    }

    const response = await fetch(fullUrl, options);
    if (!response.ok) {
        throw new Error(`HTTP error: status: ${response.status}`);
    } else if (response.status === 204) {
        return;
    }
    return await response.json();
}

export async function login(email: string, password: string): Promise<any> {
    const resp = await unauthenticatedRequest('/auth/login', 'POST', { email, password });
    if (resp.token) {
        localStorage.setItem('token', resp.token);
    }
    return resp;
}