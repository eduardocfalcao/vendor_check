import './App.css';
import { VendorCard } from './VendorCard';

function App() {
  return (
    <div data-testid="test_vendorBoard" className="VendorBoard">
      <VendorCard url='http://localhost:8000/v1/amazon-status'/>
      <VendorCard url='http://localhost:8000/v1/google-status'/>
    </div>
  );
}

export default App;
