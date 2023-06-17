const DockpanelButton = ({ icon, title, active, onClick }) => {
    return (
        <div className={`flex items-center border-[1px] border-transparent rounded-[2px] transition ${active ? 'dockpanel-button-active' : ''} cursor-pointer py-[5px] px-[10px]`} onClick={onClick}>
            <div className={`w-[60px] h-[60px] ${icon} bg-contain bg-no-repeat bg-center`} />
            <div className="text-[25px] leading-[33px] tracking-[-0.05em] ml-[20px]">{title}</div>
        </div>
    )
}

export default DockpanelButton