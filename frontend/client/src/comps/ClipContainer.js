import React, { Component } from 'react';

import YoutubeEmbed from './YoutubeEmbed';

import './ClipContainer.css';


class ClipContainer extends React.Component {

    constructor(props) {
        super(props);
    }

    state = {}


    stripURL(url) {
        return url.substring( url.indexOf("=") + 1 );
    }

    upload () {
        alert("Clip sent for upload");
    }

    render() { 

        let test = this.stripURL(this.props.link);

        return(
            <div className="clip-container">
                {this.props.vtuber.name}
                <YoutubeEmbed embedId={ this.stripURL(this.props.link) } /> 


                <div className="timestamp">
                    <div className="topdown-container">
                        <div>timestamps</div>

                        <label>Begin: <input type="text" /> </label>
                        <label>End:   <input type="text" /> </label>
                    </div>
                </div>

                <div className="description">
                    <div className="topdown-container">
                        <div>Description</div>

                        <textarea type="text" className="description-input" />

                    </div>
                </div>

                <button className="upload-button" onClick={this.upload}>Upload</button>


            </div>
        );
    }
}
 
export default ClipContainer;