import React, {useState, useEffect} from 'react';

import {Curl, getCurl} from './utils/curl';
import {DataContainer, RootContainer} from './containers/basic';
import './App.css';

interface MetadataType {
  name: string;
  content_type: string;
  size: number;
}

function App() {
  const [uploading, setUploading] = useState(false);
  const [metadata, setMetadata] = useState<MetadataType[]>([]);
  useEffect(() =>{
    const GetMetadata = async()=>{
      const url = 'http://localhost:8080/api/check-metadata';
      const headers = {
        'Authorization': 'eyJ1c2VybmFtZSI6ICJTZWVkZXIgVXNlciJ9',
      }
      try{
        var response = await getCurl(url, headers)
        setMetadata(response.data.data)
        console.log(metadata)
      } catch(err){
        alert(err)
      }
    }
    GetMetadata();
  },[uploading]);

  const onSubmit = async (e:any) => {
    const url = 'http://localhost:8080/api/upload-image';
    const headers = {
      'Authorization': 'eyJ1c2VybmFtZSI6ICJTZWVkZXIgVXNlciJ9',
    }
    setUploading(true)
    const form = new FormData();
    form.append("file", e.target.files[0], e.target.files[0].name);
    try{
      var response = await Curl(url, form, headers)
      console.log(response.data)
    } catch(err) {
      alert(err)
    }
    setUploading(false)
  }

  return (
    <RootContainer>
      <div>
        <p>Image Uploader</p>
        {
          uploading ?
          <div className="row">
            <div className="loader"></div>
            <p>Image is Uploading...</p>
          </div>
          : 
          <input type="file" onChange={onSubmit} />
        }
      </div>
      <DataContainer>
        <div className="inner-data">
          <h3>Image Metadata</h3>
          {
            metadata == null ?
            <p>Null</p>
            :
            metadata.map(item =>{
              return <div className="card">
              <p>Image Name : {item.name}</p>
              <p>Content Type : {item.content_type}</p>
              <p>Size : {item.size} KB</p>
            </div>
            })
          }
        </div>
      </DataContainer>
    </RootContainer>
  );
}

export default App;
