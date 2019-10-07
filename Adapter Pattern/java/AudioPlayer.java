/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 16:54
 **/
public class AudioPlayer {
    public void play(Audio audio){
        System.out.println(audio.getVoice());
    }

    public static void main(String[] args) {
        Media media= new Media();
        media.setTape("也不行");
        new AudioPlayer().play(new MediaAdapter().Meida2Audio(media));
    }
}
