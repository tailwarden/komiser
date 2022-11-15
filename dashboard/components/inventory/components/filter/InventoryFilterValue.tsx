import { useEffect, useState } from 'react';
import settingsService from '../../../../services/settingsService';

function InventoryFilterValue() {
  const [data, setData] = useState();
  const [error, setError] = useState(false);

  useEffect(() => {
    let mounted = true;

    settingsService.getProviders().then(res => {
      if (mounted) {
        if (res === Error) {
          setError(true);
        } else {
          setData(res);
          console.log(res);
        }
      }
    });

    return () => {
      mounted = false;
    };
  }, []);
  return <div className="flex flex-col gap-2 w-80">Hello world</div>;
}

export default InventoryFilterValue;
