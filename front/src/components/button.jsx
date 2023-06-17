import { useMediaQuery } from "../utils/useMediaQuery"

const Button = ({ children, onClick }) => {
    const isMobile = useMediaQuery(480)
    return (
        <div className={`flex items-center justify-center rounded-[3px] cursor-pointer button ${isMobile ? 'min-w-[100px] h-[30px] text-[14px]' : 'min-w-[70px] h-[25px] text-[12px]'} text-black`} onClick={onClick}>{children}</div>
    )
}

export default Button