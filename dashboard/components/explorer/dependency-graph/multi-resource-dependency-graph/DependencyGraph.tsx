/* eslint-disable react-hooks/exhaustive-deps */
import React, { memo } from 'react';
import CytoscapeComponent from 'react-cytoscapejs';
import Cytoscape from 'cytoscape';

import EmptyState from '@components/empty-state/EmptyState';

import Tooltip from '@components/tooltip/Tooltip';
import WarningIcon from '@components/icons/WarningIcon';
import DragIcon from '@components/icons/DragIcon';
import NumberInput from '@components/number-input/NumberInput';
import useInventory from '@components/inventory/hooks/useInventory/useInventory';
import InventorySidePanel from '@components/inventory/components/InventorySidePanel';
import { DependencyGraphProps } from '../hooks/useDependencyGraph';
import {
  edgeStyleConfig,
  graphLayoutConfig,
  leafStyleConfig,
  maxZoom,
  minZoom,
  nodeStyeConfig
} from '../config';
import { useDependencyGraphActions } from '../hooks/useDependencyGraphActions';

const DependencyGraph = ({ data }: DependencyGraphProps) => {
  const dataIsEmpty: boolean = data.nodes.length === 0;
  const {
    cyRef,
    cyActionHandlers,
    toggleNodeDragging,
    isNodeDraggingEnabled,
    translateXClass,
    zoomVal,
    handleZoomChange
  } = useDependencyGraphActions({ enableClickHandler: true });

  const {
    isOpen,
    closeModal,
    data: inventoryItem,
    page,
    goTo,
    tags,
    handleChange,
    addNewTag,
    removeTag,
    updateTags,
    loading,
    deleteLoading,
    bulkItems,
    updateBulkTags
  } = useInventory();

  return (
    <div className="relative h-full flex-1 bg-dependency-graph bg-[length:40px_40px]">
      {dataIsEmpty ? (
        <>
          <div className="translate-y-[201px]">
            <EmptyState
              title="No results for this filter"
              message="It seems like you have no cloud resources matching the filters you added"
              mascotPose="tablet"
            />
          </div>
        </>
      ) : (
        <>
          <CytoscapeComponent
            className="h-full w-full"
            elements={CytoscapeComponent.normalizeElements({
              nodes: data.nodes,
              edges: data.edges
            })}
            maxZoom={maxZoom}
            minZoom={minZoom}
            layout={graphLayoutConfig}
            stylesheet={[
              {
                selector: 'node',
                style: nodeStyeConfig
              },
              {
                selector: 'edge',
                style: edgeStyleConfig
              },
              {
                selector: '.leaf',
                style: leafStyleConfig
              }
            ]}
            cy={(cy: Cytoscape.Core) => {
              cyActionHandlers(cy);
              cyRef.current = cy;
            }}
          />
        </>
      )}
      <div className="absolute bottom-0 w-full">
        <div className="flex w-full flex-col items-center gap-2 sm:flex-row sm:justify-between">
          <div className="flex gap-2 overflow-visible bg-black-100 text-black-400">
            {data?.nodes?.length} Resources
            {!dataIsEmpty && (
              <div className="relative">
                <WarningIcon className="peer" height="16" width="16" />
                <Tooltip bottom="xs" align="left" width="lg">
                  Only AWS resources are currently supported on the explorer.
                </Tooltip>
              </div>
            )}
          </div>
          <div className="flex max-h-11 gap-4">
            <button
              className={`peer relative flex items-center rounded border-[1.2px] border-black-200 bg-white p-2.5 ${
                isNodeDraggingEnabled && 'border-primary'
              }`}
              onClick={toggleNodeDragging}
            >
              <DragIcon className="h-6 w-6" />
            </button>
            <Tooltip align="center" bottom="sm">
              {isNodeDraggingEnabled
                ? 'Disable node dragging'
                : 'Enable node dragging'}
            </Tooltip>

            <div className="relative w-40">
              <NumberInput
                name="zoom"
                value={zoomVal}
                action={zoomData => handleZoomChange(Number(zoomData.zoom))}
                handleValueChange={handleZoomChange} // increment or decrement input value
                step={5} // percentage change in zoom
                maxLength={3}
              />
              <span
                className={`absolute left-1/2 top-1/2 ${translateXClass} -translate-y-1/2 text-sm text-neutral-900`}
              >
                %
              </span>
            </div>
          </div>
        </div>
      </div>
      {/* Modal */}
      <InventorySidePanel
        isOpen={isOpen}
        closeModal={closeModal}
        data={inventoryItem}
        goTo={goTo}
        page={page}
        updateTags={updateTags}
        tags={tags}
        tabs={['resource details', 'tags']}
        handleChange={handleChange}
        removeTag={removeTag}
        addNewTag={addNewTag}
        loading={loading}
        deleteLoading={deleteLoading}
        bulkItems={bulkItems}
        updateBulkTags={updateBulkTags}
      />
    </div>
  );
};

export default memo(DependencyGraph);
