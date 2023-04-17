import axios from 'axios';

const URL = "http://localhost:8083"

export const findTeamName = (roles) => {
 
  return roles.includes("egress") == true ? "egress" : "mdm"
}
export const findTeam = (accessToken, roles) => {
  return roles.includes("egress") == true ? handleEgress(accessToken) : handleMDM(accessToken)
}

export const handleEgress = (token) => axios.get(`${URL}/insecure/migration/egress`).then(response=>response.data)

export const handleMDM = (token) => axios.get(`${URL}/insecure/migration/mdm`).then(response=>response.data)

export const handleLogin = (username, password,dispatcher, errorMessage) => {
  return login(username,password).then(data => {
    
    if (data.accessToken) {

      dispatcher.setAuth(data.accessToken, data.roles, data.given_name, data.family_name, data.email)
    } else {
      errorMessage(true)
      setTimeout(()=>{
        errorMessage(false)
      },3000)
    }
  })

}


export const login = (username, password) => {
  axios.defaults.headers.post['Content-Type'] ='application/json;charset=utf-8';
  axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
  return axios.get(`${URL}/login/${username}/${password}`)
  .then((response) => response.data)
  .catch((error) => error);
} 
// CORS
// export const login = (username, password) => {
//   axios.defaults.headers.post['Content-Type'] ='application/json;charset=utf-8';
//   axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
//   return axios.post(`${URL}/login`, { username, password })
//   .then((response) => response.data)
//   .catch((error) => error);
// }

const idTokenHelper = () => {
  const { token } = JSON.parse(localStorage['keyCloak'] || '{"token":"test"}');
  return token;
};
