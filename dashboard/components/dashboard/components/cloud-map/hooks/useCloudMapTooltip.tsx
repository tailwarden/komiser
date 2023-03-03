import { useState } from 'react';
import { DashboardCloudMapRegions } from './useCloudMap';

export type DashboardCloudMapTooltipProps = {
  name: string;
  label: string;
  resources: number;
  x: number;
  y: number;
};

type useCloudMapTooltipProps = {
  data: DashboardCloudMapRegions | undefined;
};

function useCloudMapTooltip({ data }: useCloudMapTooltipProps) {
  const [tooltip, setTooltip] = useState<DashboardCloudMapTooltipProps>();

  const resources = data && data.map(region => region.resources);
  const sumOfResources =
    resources && resources.reduce((resource, a) => resource + a, 0);

  return { tooltip, setTooltip, sumOfResources };
}

export default useCloudMapTooltip;
