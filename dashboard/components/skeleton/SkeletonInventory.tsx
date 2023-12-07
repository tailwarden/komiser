function SkeletonInventory() {
  const cards: number[] = Array.from(Array(6).keys());

  return (
    <div className="grid-row grid animate-pulse overflow-hidden rounded-b-lg">
      {cards.map(card => (
        <div
          key={card}
          className="flex h-[57px] items-center border-b bg-white px-6 text-sm"
        >
          <div className="flex w-full items-center gap-6">
            <div className="h-6 w-6 flex-shrink-0 rounded-full bg-cyan-200"></div>
            <div className="h-4 w-[5%] rounded-lg bg-cyan-200"></div>
            <div className="h-4 w-[20%] rounded-lg bg-cyan-200"></div>
            <div className="h-4 w-[10%] rounded-lg bg-cyan-200"></div>
            <div className="h-4 w-[30%] rounded-lg bg-cyan-200"></div>
            <div className="h-4 w-[15%] rounded-lg bg-cyan-200"></div>
            <div className="h-4 w-[5%] rounded-lg bg-cyan-200"></div>
          </div>
        </div>
      ))}
    </div>
  );
}

export default SkeletonInventory;
