import { format } from 'date-fns'; 
import './VendorCard.css';import { useVendor } from './useVendor';

export const VendorCard = (props) => {
    const { name } = props
    const { vendorData } = useVendor(name)

    return (
        <div className="VendorCard">
           <div className='Row'>
                <label className='Bold'> 
                    Url:
                </label>
                <label>
                    {vendorData.url}
                </label>
            </div>
            <div className='Row'>
                <label className='Bold'> 
                    Status Code:
                </label>
                <label>
                    {vendorData.status_code}
                </label>
            </div>
            <div className='Row'>
                <label className='Bold'> 
                    Duration:
                </label>
                <label>
                    {vendorData.duration}
                </label>
            </div>
            <div className='Row'>
                <label className='Bold'> 
                    Date:
                </label>
                <label>
                    {format(new Date(vendorData.date * 1000), `MM/dd/yyyy HH:mm:ss`)}
                </label>
            </div>
        </div>
    )
}
