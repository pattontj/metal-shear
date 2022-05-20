import React, { Component } from 'react';

import YoutubeEmbed from './YoutubeEmbed';

import './ClipContainer.css';


class ClipContainer extends React.Component {

    constructor(props) {
        super(props);
    }

    state = {}


    upload () {
        alert("Clip sent for upload");
    }

    render() { 
        return(
            <div className="clip-container">
                {this.props.name}
                <YoutubeEmbed embedId="reJgFRAGf5o" /> 


                <button className="upload-button" onClick={this.upload}>Upload</button>

            </div>
        );
    }
}
 
export default ClipContainer;