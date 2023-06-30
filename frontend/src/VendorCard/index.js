import { useEffect, useState } from 'react';
import { format } from 'date-fns'; 
import './VendorCard.css';


export const VendorCard = (props) => {
    const { url } = props
    const [vendor, setVendor] = useState({
        url: '',
        status_code: 0,
        duration: 0,
        date: 0,
    })
    useEffect(() => {
        fetch(url)
            .then((res) => res.json())
            .then((data) => {
                setVendor(data);
         })
         .catch((err) => {
            console.log(err);
         });
    }, [url])

    return (
        <div className="VendorCard">
           <div className='Row'>
                <label className='Bold'> 
                    Url:
                </label>
                <label>
                    {vendor.url}
                </label>
            </div>
            <div className='Row'>
                <label className='Bold'> 
                    Status Code:
                </label>
                <label>
                    {vendor.status_code}
                </label>
            </div>
            <div className='Row'>
                <label className='Bold'> 
                    Duration:
                </label>
                <label>
                    {vendor.duration}
                </label>
            </div>
            <div className='Row'>
                <label className='Bold'> 
                    Date:
                </label>
                <label>
                    {format(new Date(vendor.date * 1000), `MM/dd/yyyy HH:mm:ss`)}
                </label>
            </div>
        </div>
    )
}
