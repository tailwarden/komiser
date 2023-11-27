import { memo } from 'react';
import { DashboardCloudMapTooltipProps } from './hooks/useCloudMapTooltip';
import LinkIcon from '../../../icons/LinkIcon';

type DashboardCloudMapProps = {
  tooltip: DashboardCloudMapTooltipProps | undefined;
  sumOfResources: number | undefined;
};

function DashboardCloudMap({
  tooltip,
  sumOfResources
}: DashboardCloudMapProps) {
  return (
    <>
      {tooltip && sumOfResources && (
        <div
          className="absolute z-20 flex animate-fade-in flex-col gap-2 rounded-lg bg-[#013220] px-3 py-2 text-xs text-gray-500 opacity-0"
          style={{
            top: `${tooltip.y - 60}px`,
            left: `${tooltip.x + 10}px`
          }}
        >
          <div className="-mx-3 flex items-center gap-2 border-b border-white/30 px-3 pb-2">
            <div
              className={`h-2 w-2 rounded-full ${
                tooltip.resources === 0 ? 'bg-gray-500' : 'bg-blue-500'
              }`}
            ></div>
            <span className="font-medium text-white">{tooltip.name} </span>
            <span className="font-medium text-white"> - {tooltip.label}</span>
          </div>

          <span>
            Active resources:{' '}
            <span className="font-medium text-white">{tooltip.resources}</span>
          </span>
          <span>
            Percentage:{' '}
            <span className="font-medium text-white">{`${(
              (tooltip.resources / sumOfResources) *
              100
            ).toFixed(1)}%`}</span>
          </span>
          <div className="-mx-3 flex items-center gap-2 border-t border-white/30 px-3 pb-1 pt-2">
            <span className="text-white">Click to discover the resources</span>
            <LinkIcon className="w-[20px]" />
          </div>
        </div>
      )}
    </>
  );
}

export default memo(DashboardCloudMap);
