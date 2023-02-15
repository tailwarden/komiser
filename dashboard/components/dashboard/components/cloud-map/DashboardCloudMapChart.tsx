import { memo } from 'react';
import {
  ComposableMap,
  Geographies,
  Geography,
  Marker
} from 'react-simple-maps';
import { Regions } from './useCloudMap';
import { CloudMapTooltip } from './useCloudMapTooltip';

type DashboardCloudMapChartProps = {
  regions: Regions | undefined;
  setTooltip: (tooltip: CloudMapTooltip | undefined) => void;
};

function DashboardCloudMapChart({
  regions,
  setTooltip
}: DashboardCloudMapChartProps) {
  const geoUrl = '/data/map/countries.json';

  const regionsAscendingByNumberOfResources =
    regions && regions.sort((a, b) => a.resources - b.resources);

  return (
    <div className="relative">
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
                fill={'#DAD3E2'}
                stroke={'white'}
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
            <Marker
              key={idx}
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
                  r={10}
                  fill="#56BA5B"
                  className="pointer-events-none animate-wide-pulse"
                />
              )}
              <circle
                r={10}
                fill={region.resources === 0 ? '#978EA1' : '#56BA5B'}
                stroke={'white'}
                strokeWidth={1.5}
              />
            </Marker>
          ))}
      </ComposableMap>
    </div>
  );
}

export default memo(DashboardCloudMapChart);
