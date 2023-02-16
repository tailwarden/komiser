import { useState } from 'react';

export type DashboardCloudMapTooltipProps = {
  name: string;
  label: string;
  resources: number;
  x: number;
  y: number;
};

function useCloudMapTooltip() {
  const [tooltip, setTooltip] = useState<DashboardCloudMapTooltipProps>();

  return { tooltip, setTooltip };
}

export default useCloudMapTooltip;
