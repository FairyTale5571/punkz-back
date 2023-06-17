const TaskbarButton = ({ icon, onClick }) => {
    return (
        <div className={`relative taskbar-button button-radial-gradient w-[60px] h-[39px] cursor-pointer button-sep-w-gradient button-sep-b-gradient after:content-[''] after:absolute after:h-full after:w-[1px] after:top-0 after:right-[1px] before:content-[''] before:absolute before:h-full before:w-[1px] before:top-0 before:right-0 rounded-[2px]`} onClick={onClick}>
            <div className={`w-full h-full ${icon} bg-center bg-no-repeat button-sep-w-gradient button-sep-b-gradient after:content-[''] after:absolute after:h-full after:w-[1px] after:left-0 after:top-0 before:content-[''] before:absolute before:h-full before:w-[1px] before:top-0 before:left-[1px]`} />
            <div className="hoverable button-taskbar-hover-gradient" />
        </div>
    )
}

export default TaskbarButton