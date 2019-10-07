/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 15:59
 **/
public class GothamBuilder {
    public City buildGotham() {
        City city = new City();
        city.setHero(new Batman());
        city.setName("Gotham");
        return city;
    }
}
