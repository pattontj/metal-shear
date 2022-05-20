import React, { Component } from 'react';

import ClipContainer from './comps/ClipContainer';

class Home extends React.Component {

    state = 
    {
        vtubers: [],
    }

    componentDidMount() {

        fetch("http://localhost:8080/api/vtubers")
            .then( response => response.json() )
            .then ( 
               data => this.setState( {vtubers: data} )
            )
            .catch(err => console.log("response failed, ", err));

    }

    render() { 
        return(
        <div style={{alignContent: "center"}} >
            <h1>Hello, Sailor!</h1> 
            
        {/* <li key={vtuber.toString()}>{vtuber.name}</li>*/}

         
            { this.state.vtubers.map( (vtuber) => <ClipContainer key={vtuber.id} name={vtuber.name} /> )}
          
        </div>
        )
    }
}
 
export default Home; <h1>Hello, Sailor!</h1> 