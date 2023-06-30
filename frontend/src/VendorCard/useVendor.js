import { useEffect, useState } from 'react';

const MINUTE = 60000;

const VENDOR_URL = {
    'amazon': 'http://localhost:8000/v1/amazon-status',
    'google': 'http://localhost:8000/v1/google-status',
}

export const useVendor = (vendor) => {
    const url = VENDOR_URL[vendor]
    const [vendorData, setVendorData] = useState({
        url: url,
        status_code: 0,
        duration: 0,
        date: 0,
    })

    const fetchVendor = (url) => {
        fetch(url)
            .then((res) => res.json())
            .then((data) => {
                setVendorData(data);
            })
            .catch((err) => {
                console.log(err);
            });
    } 

    useEffect(() => {
        fetchVendor(url);
        const interval = setInterval(() => {
            fetchVendor(url)
        }, MINUTE);

         return () => clearInterval(interval)
    }, [url])

    return { vendorData }
}
