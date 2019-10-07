/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 15:54
 **/
public class City {
    protected String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void heroSpeech(Hero hero){
        System.out.println("这里是"+name);
        hero.sayHello();
    }

}
