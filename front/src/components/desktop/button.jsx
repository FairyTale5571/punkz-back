const DesktopButton = ({ icon, title, className, hoverable, hover, onClick }) => {
    return (
        <div className={`absolute z-10 ${hover ? 'desktop-button-hover' : ''} ${hoverable ? 'cursor-pointer desktop-button' : ''} h-[133px] min-w-[110px] flex flex-col justify-center ${className}`} onClick={onClick}>
            <div className={`h-[85px] bg-center bg-no-repeat bg-contain ${icon}`} />
            <div className="flex justify-center mt-[5px] text-[25px] tracking-[-0.05em] leading-[33px]">{title}</div>
        </div>
    )
}

export default DesktopButton