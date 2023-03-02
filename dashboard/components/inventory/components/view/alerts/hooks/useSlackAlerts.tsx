import { useEffect, useState } from 'react';
import settingsService from '../../../../../../services/settingsService';

function useSlackAlerts() {
  const [loading, setLoading] = useState(false);
  const [hasSlack, setHasSlack] = useState(false);
  const [error, setError] = useState(false);

  function fetchSlackStatus() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getSlackIntegration().then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setHasSlack(res.enabled);
      }
    });
  }

  useEffect(() => {
    if (!hasSlack) {
      fetchSlackStatus();
    }

    if (hasSlack) {
      console.log('olar', hasSlack);
    }
  }, [hasSlack]);

  return { loading, hasSlack, error };
}

export default useSlackAlerts;
