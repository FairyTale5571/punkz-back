import { useEffect, useState } from "react";

const Clock = () => {
    const [date, setDate] = useState(new Date());

    useEffect(() => {
      const timerID = setInterval(() => tick(), 1000);
      return function cleanup() {
        clearInterval(timerID);
      };
    });
  
    function tick() {
      setDate(new Date());
    }
  
    const timeOptions = { hour: 'numeric', minute: '2-digit', hour12: true };
    const time = date.toLocaleTimeString([], timeOptions);
  
    const day = date.getDate();
    const month = date.getMonth() + 1;
    const year = date.getFullYear();
    const currentDate = `${month}/${day}/${year}`;
  
    return (
      <div className="flex flex-col items-center text-[12px] text-white drop-shadow-text clock-shadow tracking-[-0.005em] leading-[135.51%] justify-center whitespace-nowrap">
        <span>{time}</span>
        <span>{currentDate}</span>
      </div>
    );
}

export default Clock