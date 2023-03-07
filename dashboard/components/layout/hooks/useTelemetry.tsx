import { useEffect, useState } from 'react';
import settingsService from '../../../services/settingsService';

type Telemetry = {
  telemetry_enabled: boolean;
};

function useTelemetry() {
  const [telemetry, setTelemetry] = useState<Telemetry>();

  useEffect(() => {
    settingsService.getTelemetry().then(res => {
      if (res === Error) {
        throw new Error(
          'Server could not be found. Refer to the logs and try again.'
        );
      } else {
        setTelemetry(res);
      }
    });
  }, []);
  return { telemetry };
}

export default useTelemetry;
