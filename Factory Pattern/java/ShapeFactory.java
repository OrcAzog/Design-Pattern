/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 11:51
 **/
public class ShapeFactory {
    public Shape getShape(String name){
        if(name.equals("circle"))return new Circle();
        if(name.equals("rectangle"))return new Rectangle();
        if(name.equals("square"))return new Square();
        else return null;
    }
}
