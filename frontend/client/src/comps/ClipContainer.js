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

    postData(vid_id, start, end) {
        console.log(vid_id, start, end);
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json;charset=UTF-8', 'Accept': 'application/json',},
            body: JSON.stringify({video_id: vid_id, start:start, end:end})
        };
        fetch('http://192.168.0.13:5050/api/v1/download', requestOptions)
            .then(data => console.log(data.toString()));
    }

    render() { 

        // let test = this.stripURL(this.props.link);

        return(
            <div className="clip-container">
                <YoutubeEmbed embedId={ this.props.link } start={this.props.start_s} end={this.props.end_s} /> 


                <div className="timestamp">
                    <div className="topdown-container">
                        <div>timestamps</div>

                        <label>Begin: <input type="number" id="start" defaultValue={this.props.start_s}/> </label>
                        <label>End:   <input type="number" id="end" defaultValue={this.props.end_s}/> </label>
                    </div>
                </div>

                <div className="description">
                    <div className="topdown-container">
                        <div>Title</div>
                        <textarea type="text" className="title-input" />
                        <div>Description</div>
                        <textarea type="text" className="description-input" />

                    </div>
                </div>

                <button className="upload-button" onClick={e => this.postData(this.props.link, this.props.start_s, this.props.end_s)}>Upload</button>


            </div>
        );
    }
}
export default ClipContainer;