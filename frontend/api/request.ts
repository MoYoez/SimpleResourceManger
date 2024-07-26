import axios from 'axios'

interface BaseResponse {
    status: boolean;
    message: string;
}


axios.defaults.baseURL = 'http://localhost:8080/backend';
axios.defaults.headers.common['Authorization'] = ReadAuth();


export function SetAuth(AuthToken: string) {
    var date = new Date();
    date.setTime(date.getTime() + (24 * 60 * 60 * 1000));
    var expires = "expires=" + date.toUTCString();
    document.cookie = "Authorization="+AuthToken+ ";"+ expires + ";+ path=/";
    axios.defaults.headers.common['Authorization'] = ReadAuth();
}

export function ReadAuth() {
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim().split('=');
        if (cookie[0] === 'Authorization') {
            return cookie[1];
        }
    }

    return '';
}




export type { BaseResponse };
