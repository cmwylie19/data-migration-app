import * as React from 'react'
import TopNav from '../TopNav/TopNav'
import NameCard from "../NameCard/NameCard"
import { findTeam, findTeamName } from "../../API"

 const Landing = ({authReducer}) => {
    const [team, setTeam] = React.useState([])
    const {roles,given_name } = authReducer.auth
    React.useEffect(()=>{
        findTeam(authReducer.auth.accessToken, authReducer.auth.roles).then(data=>setTeam(data))
    },[])

    return (
        <>
        <TopNav given_name={given_name} authReducer={authReducer} />
        {Array.isArray(team.members) && team.members.map(member=>
            <NameCard team={findTeamName(roles)} first={member.first} last={member.last} />
        )}
  
        </>
    )
}
export default Landing