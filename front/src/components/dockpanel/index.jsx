import { Transition } from "@headlessui/react"
import { Fragment } from "react"
import DockpanelButton from "./button"

const DockPanel = ({ open, isMobile, actions }) => {
    const { connectWallet, setConnectWallet, about, setAbout, auth } = actions
    return (
        <Transition
            show={open}
            enter="transition duration-500 ease-out"
            enterFrom="transform translate-y-[500px] opacity-0"
            enterTo="transform translate-y-[0px] opacity-500"
            leave="transition duration-300 ease-out"
            leaveFrom="transform translate-y-[0px] opacity-100"
            leaveTo="transform translate-y-[500px] opacity-0"
            className={`h-full z-20 ${isMobile ? '' : 'absolute'}`}
            as={'div'}
        >
            <div className={`absolute dockpanel ${isMobile ? 'w-full h-full bottom-0' : 'h-[466px] w-[414px] bottom-[41px]'} left-0`}>
                <div className="relative flex w-full h-full p-[8px]">
                    <div className={`${isMobile ? 'max-w-[360px] w-full' : 'w-[245px]'} h-full bg-white border-[#9D9D9D] border-[1px] border-solid rounded-tr-[2px] p-[8px]`}>
                        {!(isMobile) ? (
                            <Fragment>
                                <div className="flex items-center w-full">
                                    <div className="bg-help bg-center bg-no-repeat bg-contain w-[30px] h-[30px]" />
                                    <div className="ml-[8px] text-[14px] leading-[19px] tracking-[-0.05em]">About Punkz</div>
                                </div>
                                <div className="mt-[22px] text-[20px] leading-[27px] tracking-[-0.05em] w-[210px] text-[#000000E4]">
                                    PUNKZ is an ecosystem of NFT utilities. The main utility is Omnistep, it is a platform to comfortably manage your project and to keep all your data and tasks in a simple way. The main utility of our platform is that we become a guarantor of work done and money paid.
                                </div>
                            </Fragment>
                        ) : (
                            <div className="flex flex-col w-full h-full py-[12px] px-[6px] gap-[40px]">
                                <DockpanelButton icon={'bg-computer'} title={'Connect Wallet'} active={connectWallet} onClick={() => setConnectWallet(prevState => !prevState)} />
                                <DockpanelButton icon={'bg-help'} title={'About PUNKZ'} active={about} onClick={() => setAbout(prevState => !prevState)} />
                                <DockpanelButton icon={'bg-recycle'} title={'Recycle bin'} onClick={() => window.open('https://punkz.gitbook.io/punkz/', '_blank')} />
                                <DockpanelButton icon={'bg-twitter-logo'} title={'Twitter'} onClick={() => window.open('https://twitter.com/wepunkz', '_blank')} />
                                <DockpanelButton icon={'bg-discord-logo'} title={'Discord'} onClick={() => window.open('https://discord.gg/wepunkz', '_blank')} />
                            </div>
                        )}
                    </div>
                    {!(isMobile) && (
                        <Fragment>
                            <div className="relative h-full w-full ml-[14px]">
                                <div className="m-auto mt-[-40px] w-[64px] h-[64px] rounded-[8.55556px] border-[0.5px] dockpanel-avatar border-solid border-[#00000059] p-[6px]">
                                    <div className="h-full w-full dockpanel-avatar-square p-[2px]">
                                        <div className="h-full w-full dockpanel-avatar-item">
                                            <div className={`h-full w-full bg-avatar bg-no-repeat bg-contain bg-center`} style={auth?.avatar ? {backgroundImage: `url(${auth.avatar})`} : {}}></div>
                                        </div>
                                    </div>
                                </div>
                                <div className="flex flex-col text-[#FFFFFFE4] mt-[30px] ml-[6px]">
                                    <div className="dockpanel-text text-[13px] leading-[17px] tracking-[-0.05em] cursor-pointer">Wikipedia</div>
                                    <div className="mt-[17px] dockpanel-text text-[13px] leading-[17px] tracking-[-0.05em] cursor-pointer">Google</div>
                                    <div className="mt-[17px] dockpanel-text text-[13px] leading-[17px] tracking-[-0.05em] cursor-pointer">Rick Roll</div>
                                    <div className="mt-[17px] dockpanel-text text-[13px] leading-[17px] tracking-[-0.05em] cursor-pointer">Omnistep Twitter</div>
                                </div>
                            </div>
                        </Fragment>
                    )}
                </div>
                <div className="absolute w-full h-[200px] left-0 bottom-0 z-0 glare-b-r" />
            </div>
        </Transition>
    )
}

export default DockPanel