import { useState } from 'react';

function useTelemetry() {
  const [telemetry, setTelemetry] = useState(false);
  return { telemetry };
}

export default useTelemetry;
