/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 18:07
 **/
public class ObsceneLanguageFilter implements LanguageFilter {
    @Override
    public Language filter(Language language) {
        StringBuilder stringBuilder = new StringBuilder();
        if (language.getString().contains("fuck")) {
            String[] strings = language.getString().split("fuck");
            for (String str : strings) {
                stringBuilder.append(str);
            }
        }
        language.setString(stringBuilder.toString());
        return language;
    }
}
