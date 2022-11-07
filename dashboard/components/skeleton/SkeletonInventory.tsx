function SkeletonInventory() {
  const cards: number[] = Array.from(Array(6).keys());

  return (
    <div className="grid grid-row animate-pulse rounded-b-lg overflow-hidden">
      {cards.map(card => (
        <div
          key={card}
          className="flex items-center h-[57px] px-6 text-sm bg-white border-b"
        >
          <div className="w-full flex items-center gap-6">
            <div className="flex-shrink-0 w-6 h-6 rounded-full bg-komiser-200/50"></div>
            <div className="w-[5%] h-4 bg-komiser-200/50 rounded-lg"></div>
            <div className="w-[20%] h-4 bg-komiser-200/50 rounded-lg"></div>
            <div className="w-[10%] h-4 bg-komiser-200/50 rounded-lg"></div>
            <div className="w-[30%] h-4 bg-komiser-200/50 rounded-lg"></div>
            <div className="w-[15%] h-4 bg-komiser-200/50 rounded-lg"></div>
            <div className="w-[5%] h-4 bg-komiser-200/50 rounded-lg"></div>
          </div>
        </div>
      ))}
    </div>
  );
}

export default SkeletonInventory;
