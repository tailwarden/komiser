import { memo } from 'react';
import {
  ComposableMap,
  Geographies,
  Geography,
  Marker
} from 'react-simple-maps';
import Link from 'next/link';
import { DashboardCloudMapRegions } from './hooks/useCloudMap';
import { DashboardCloudMapTooltipProps } from './hooks/useCloudMapTooltip';

type DashboardCloudMapChartProps = {
  regions: DashboardCloudMapRegions | undefined;
  setTooltip: (tooltip: DashboardCloudMapTooltipProps | undefined) => void;
  isOpen: boolean;
};

function DashboardCloudMapChart({
  regions,
  setTooltip,
  isOpen
}: DashboardCloudMapChartProps) {
  const geoUrl = '/data/map/countries.json';

  const regionsAscendingByNumberOfResources =
    regions && regions.sort((a, b) => a.resources - b.resources);

  return (
    <ComposableMap
      projection="geoNaturalEarth1"
      height={482}
      width={820}
      projectionConfig={{
        center: [14, 0]
      }}
    >
      <Geographies geography={geoUrl}>
        {({ geographies }) =>
          geographies.map(geo => (
            <Geography
              tabIndex={-1}
              key={geo.rsmKey}
              geography={geo}
              fill="#F4F9F9"
              stroke="#95A3A3"
              strokeWidth={0.35}
              style={{
                default: { outline: 'none' },
                hover: { outline: 'none' },
                pressed: { outline: 'none' }
              }}
            />
          ))
        }
      </Geographies>
      {regionsAscendingByNumberOfResources &&
        regionsAscendingByNumberOfResources.map((region, idx) => (
          <Link key={idx} href={`/inventory?region:IS:${region.label}`}>
            <Marker
              coordinates={[Number(region.longitude), Number(region.latitude)]}
              onMouseEnter={e =>
                setTooltip({
                  name: region.name,
                  label: region.label,
                  resources: region.resources,
                  x: e.pageX,
                  y: e.pageY
                })
              }
              onMouseLeave={() => setTooltip(undefined)}
            >
              {region.resources > 0 && (
                <circle
                  r={isOpen ? 6 : 8.5}
                  fill="#387BEB"
                  className="pointer-events-none animate-wide-pulse"
                />
              )}
              <circle
                r={isOpen ? 6 : 8.5}
                fill={region.resources === 0 ? '#95A3A3' : '#387BEB'}
                stroke="white"
                strokeWidth={1.5}
              />
            </Marker>
          </Link>
        ))}
    </ComposableMap>
  );
}

export default memo(DashboardCloudMapChart);
