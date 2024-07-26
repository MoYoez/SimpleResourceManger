import axios from "axios";

export interface LoginRequest {
    user_name: string;
    user_password:  string;
}

export function Login(form:LoginRequest) {
    return axios.post("/annoy/login", form)
}
