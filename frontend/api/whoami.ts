import axios from "axios";

export function getWhoami() {
    return axios.get("/normal/whoami");
}
