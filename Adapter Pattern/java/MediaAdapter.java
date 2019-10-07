/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 16:58
 **/
public class MediaAdapter {
    public Audio Meida2Audio(Media media){
        Audio audio=new Audio();
        audio.setVoice(media.getTape());
        return audio;
    }
}
