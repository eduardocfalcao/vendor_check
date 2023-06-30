import './App.css';
import { VendorCard } from './VendorCard';

function App() {
  return (
    <div data-testid="test_vendorBoard" className="VendorBoard">
      <VendorCard name='amazon'/>
      <VendorCard name='google'/>
    </div>
  );
}

export default App;
