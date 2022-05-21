import React, { Component } from 'react';

import ClipContainer from './comps/ClipContainer';

class Home extends React.Component {

    state = 
    {
        clips: [],
    }

    componentDidMount() {

        fetch("http://localhost:8080/api/clips")
            .then( response => response.json() )
            .then ( 
               data => this.setState( {clips: data} )
            )
            .catch(err => console.log("response failed, ", err));

    }

    render() { 
        console.log(this.state.clips)
        return(
        <div style={{alignContent: "center"}} >
            <h1>Hello, Sailor!</h1> 
            
        {/* <li key={vtuber.toString()}>{vtuber.name}</li>*/}

         
            { this.state.clips.map( (clip) => <ClipContainer key={clip.id} vtuber={clip.vtuber} link={clip.link} /> )}
          
        </div>
        )
    }
}
 
export default Home; <h1>Hello, Sailor!</h1> 