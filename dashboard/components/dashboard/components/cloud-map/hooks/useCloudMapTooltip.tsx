import { useState } from 'react';

export type CloudMapTooltip = {
  name: string;
  label: string;
  resources: number;
  x: number;
  y: number;
};

function useCloudMapTooltip() {
  const [tooltip, setTooltip] = useState<CloudMapTooltip>();

  return { tooltip, setTooltip };
}

export default useCloudMapTooltip;
